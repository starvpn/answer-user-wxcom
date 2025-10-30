package wechat

import (
	"crypto/rand"
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/apache/answer/plugin"
)

//go:embed i18n/en_US.yaml
var enUSTranslationFile []byte

//go:embed i18n/zh_CN.yaml
var zhCNTranslationFile []byte

// Connector WeChat OAuth2 connector plugin
type Connector struct {
	Config *Config
}

// Config WeChat connector configuration
type Config struct {
	AppID       string `json:"app_id"`
	AppSecret   string `json:"app_secret"`
	DisplayName string `json:"display_name"`
}

// WeChatUserInfo WeChat user information response
type WeChatUserInfo struct {
	OpenID     string   `json:"openid"`
	UnionID    string   `json:"unionid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	ErrCode    int      `json:"errcode"`
	ErrMsg     string   `json:"errmsg"`
}

// WeChatAccessToken WeChat access token response
type WeChatAccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
	UnionID      string `json:"unionid"`
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
}

func init() {
	plugin.Register(&Connector{
		Config: &Config{},
	})
}

// generateState generates a random state string for CSRF protection
func generateState() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

// Info implements plugin.Base
func (c *Connector) Info() plugin.Info {
	return plugin.Info{
		Name:        plugin.MakeTranslator(InfoName),
		SlugName:    "wechat_connector",
		Description: plugin.MakeTranslator(InfoDescription),
		Author:      "Answer Team",
		Version:     "1.0.0",
		Link:        "https://github.com/starvpn/answer-user-wxcom",
	}
}

// ConnectorLogoSVG implements plugin.Connector
func (c *Connector) ConnectorLogoSVG() string {
	return `<svg width="24" height="24" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
		<path fill="#09BB07" d="M9.5 4.5c-3.59 0-6.5 2.46-6.5 5.5 0 1.78.94 3.37 2.44 4.38-.14.42-.42 1.27-.49 1.47-.08.28.1.27.22.2.09-.05 1.33-.88 1.54-1.02C7.39 15.03 8.42 15.5 9.5 15.5c3.59 0 6.5-2.46 6.5-5.5s-2.91-5.5-6.5-5.5z"/>
		<circle fill="#09BB07" cx="7" cy="10" r=".75"/>
		<circle fill="#09BB07" cx="9.5" cy="10" r=".75"/>
		<circle fill="#09BB07" cx="12" cy="10" r=".75"/>
		<path fill="#09BB07" d="M17.5 11c-.19 0-.37.01-.56.03.33-.59.56-1.25.56-1.98 0-3.59-3.36-6.5-7.5-6.5-3.78 0-6.89 2.41-7.42 5.5C1.19 8.69 0 10.22 0 12c0 1.5.81 2.86 2.1 3.73-.12.35-.36 1.08-.42 1.25-.08.24.08.23.19.17.08-.04 1.13-.75 1.31-.87.58.16 1.2.25 1.82.25.19 0 .37-.01.56-.03-.33.59-.56 1.25-.56 1.98 0 3.59 3.36 6.5 7.5 6.5 1.36 0 2.63-.32 3.73-.87.18.12 1.23.83 1.31.87.11.06.27.07.19-.17-.06-.17-.3-.9-.42-1.25C19.19 21.86 20 20.5 20 19c0-2.76-2.69-5-6-5zm-3 8c-.93 0-1.83-.19-2.63-.52-.17.11-1.08.71-1.08.71.05-.14.26-.78.34-1.03C9.73 17.4 9 16.28 9 15c0-2.48 2.46-4.5 5.5-4.5s5.5 2.02 5.5 4.5-2.46 4.5-5.5 4.5z"/>
		<circle fill="#09BB07" cx="13" cy="15" r=".65"/>
		<circle fill="#09BB07" cx="15" cy="15" r=".65"/>
		<circle fill="#09BB07" cx="17" cy="15" r=".65"/>
	</svg>`
}

// ConnectorName implements plugin.Connector
func (c *Connector) ConnectorName() plugin.Translator {
	return plugin.MakeTranslator(ConnectorName)
}

// ConnectorSlugName implements plugin.Connector
func (c *Connector) ConnectorSlugName() string {
	return "wechat"
}

// ConnectorSender implements plugin.Connector
// It redirects user to WeChat OAuth authorization page
func (c *Connector) ConnectorSender(ctx *plugin.GinContext, receiverURL string) (redirectURL string) {
	state := generateState()
	ctx.SetCookie("wechat_state", state, 600, "/", "", false, true)

	// WeChat OAuth2 authorization URL
	authURL := "https://open.weixin.qq.com/connect/qrconnect"
	params := url.Values{}
	params.Add("appid", c.Config.AppID)
	params.Add("redirect_uri", receiverURL)
	params.Add("response_type", "code")
	params.Add("scope", "snsapi_login")
	params.Add("state", state)

	return fmt.Sprintf("%s?%s#wechat_redirect", authURL, params.Encode())
}

// ConnectorReceiver implements plugin.Connector
// It handles the callback from WeChat and retrieves user information
func (c *Connector) ConnectorReceiver(ctx *plugin.GinContext, receiverURL string) (userInfo plugin.ExternalLoginUserInfo, err error) {
	// Verify state to prevent CSRF
	state := ctx.Query("state")
	savedState, _ := ctx.Cookie("wechat_state")
	if state == "" || state != savedState {
		return userInfo, fmt.Errorf("invalid state parameter")
	}

	// Get authorization code
	code := ctx.Query("code")
	if code == "" {
		return userInfo, fmt.Errorf("code not found")
	}

	// Exchange code for access token
	token, err := c.getAccessToken(code)
	if err != nil {
		return userInfo, err
	}

	// Get user information
	wechatUser, err := c.getUserInfo(token.AccessToken, token.OpenID)
	if err != nil {
		return userInfo, err
	}

	// Map WeChat user info to Answer external login user info
	// Use UnionID if available, otherwise use OpenID
	externalID := token.UnionID
	if externalID == "" {
		externalID = token.OpenID
	}

	// Generate safe username from OpenID
	username := "wechat_" + token.OpenID
	if len(token.OpenID) > 8 {
		username = "wechat_" + token.OpenID[:8]
	}

	userInfo = plugin.ExternalLoginUserInfo{
		ExternalID:  externalID,
		DisplayName: wechatUser.Nickname,
		Username:    username,
		Email:       "", // WeChat doesn't provide email
		Avatar:      wechatUser.HeadImgURL,
		MetaInfo:    fmt.Sprintf(`{"openid":"%s","unionid":"%s"}`, token.OpenID, token.UnionID),
	}

	return userInfo, nil
}

// getAccessToken exchanges authorization code for access token
func (c *Connector) getAccessToken(code string) (*WeChatAccessToken, error) {
	apiURL := "https://api.weixin.qq.com/sns/oauth2/access_token"
	params := url.Values{}
	params.Add("appid", c.Config.AppID)
	params.Add("secret", c.Config.AppSecret)
	params.Add("code", code)
	params.Add("grant_type", "authorization_code")

	resp, err := http.Get(fmt.Sprintf("%s?%s", apiURL, params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var token WeChatAccessToken
	if err := json.Unmarshal(body, &token); err != nil {
		return nil, fmt.Errorf("failed to parse token response: %w", err)
	}

	if token.ErrCode != 0 {
		return nil, fmt.Errorf("WeChat API error: %s (code: %d)", token.ErrMsg, token.ErrCode)
	}

	return &token, nil
}

// getUserInfo retrieves user information from WeChat API
func (c *Connector) getUserInfo(accessToken, openID string) (*WeChatUserInfo, error) {
	apiURL := "https://api.weixin.qq.com/sns/userinfo"
	params := url.Values{}
	params.Add("access_token", accessToken)
	params.Add("openid", openID)
	params.Add("lang", "zh_CN") // Request Chinese language

	resp, err := http.Get(fmt.Sprintf("%s?%s", apiURL, params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var userInfo WeChatUserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, fmt.Errorf("failed to parse user info: %w", err)
	}

	// Check for WeChat API errors
	if userInfo.ErrCode != 0 {
		return nil, fmt.Errorf("WeChat API error: %s (code: %d)", userInfo.ErrMsg, userInfo.ErrCode)
	}

	return &userInfo, nil
}

// ConfigFields implements plugin.Config
func (c *Connector) ConfigFields() []plugin.ConfigField {
	return []plugin.ConfigField{
		{
			Name:        "app_id",
			Type:        plugin.ConfigTypeInput,
			Title:       plugin.MakeTranslator(ConfigAppIDTitle),
			Description: plugin.MakeTranslator(ConfigAppIDDescription),
			Required:    true,
			UIOptions: plugin.ConfigFieldUIOptions{
				InputType: plugin.InputTypeText,
			},
			Value: c.Config.AppID,
		},
		{
			Name:        "app_secret",
			Type:        plugin.ConfigTypeInput,
			Title:       plugin.MakeTranslator(ConfigAppSecretTitle),
			Description: plugin.MakeTranslator(ConfigAppSecretDescription),
			Required:    true,
			UIOptions: plugin.ConfigFieldUIOptions{
				InputType: plugin.InputTypePassword,
			},
			Value: c.Config.AppSecret,
		},
		{
			Name:        "display_name",
			Type:        plugin.ConfigTypeInput,
			Title:       plugin.MakeTranslator(ConfigDisplayNameTitle),
			Description: plugin.MakeTranslator(ConfigDisplayNameDescription),
			Required:    false,
			UIOptions: plugin.ConfigFieldUIOptions{
				InputType: plugin.InputTypeText,
			},
			Value: c.Config.DisplayName,
		},
	}
}

// ConfigReceiver implements plugin.Config
func (c *Connector) ConfigReceiver(config []byte) error {
	conf := &Config{}
	if err := json.Unmarshal(config, conf); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	c.Config = conf
	return nil
}

// TranslateConfig implements plugin.Translator
func (c *Connector) TranslateConfig(language string) string {
	switch language {
	case "zh_CN", "zh-CN", "zh":
		return string(zhCNTranslationFile)
	default:
		return string(enUSTranslationFile)
	}
}

// TranslateConfigList implements plugin.Translator
func (c *Connector) TranslateConfigList() []string {
	return []string{"zh_CN"}
}
