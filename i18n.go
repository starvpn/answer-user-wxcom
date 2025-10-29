package wechat

import "github.com/apache/answer/plugin"

// i18n constants
const (
	InfoName        = "plugin.wechat_connector.name"
	InfoDescription = "plugin.wechat_connector.description"

	ConnectorName = "plugin.wechat_connector.name"

	ConfigAppIDTitle             = "plugin.wechat_connector.config.app_id.title"
	ConfigAppIDDescription       = "plugin.wechat_connector.config.app_id.description"
	ConfigAppSecretTitle         = "plugin.wechat_connector.config.app_secret.title"
	ConfigAppSecretDescription   = "plugin.wechat_connector.config.app_secret.description"
	ConfigDisplayNameTitle       = "plugin.wechat_connector.config.display_name.title"
	ConfigDisplayNameDescription = "plugin.wechat_connector.config.display_name.description"
)

var i18n = plugin.TranslatorMap{
	InfoName: {
		plugin.LanguageEnglish: "WeChat",
		plugin.LanguageChinese: "微信",
	},
	InfoDescription: {
		plugin.LanguageEnglish: "Enable users to log in with their WeChat account",
		plugin.LanguageChinese: "允许用户使用微信账号登录",
	},
	ConfigAppIDTitle: {
		plugin.LanguageEnglish: "AppID",
		plugin.LanguageChinese: "AppID",
	},
	ConfigAppIDDescription: {
		plugin.LanguageEnglish: "WeChat Open Platform Application ID",
		plugin.LanguageChinese: "微信开放平台应用 ID",
	},
	ConfigAppSecretTitle: {
		plugin.LanguageEnglish: "AppSecret",
		plugin.LanguageChinese: "AppSecret",
	},
	ConfigAppSecretDescription: {
		plugin.LanguageEnglish: "WeChat Open Platform Application Secret",
		plugin.LanguageChinese: "微信开放平台应用密钥",
	},
	ConfigDisplayNameTitle: {
		plugin.LanguageEnglish: "Display Name",
		plugin.LanguageChinese: "显示名称",
	},
	ConfigDisplayNameDescription: {
		plugin.LanguageEnglish: "Custom display name for WeChat login button",
		plugin.LanguageChinese: "微信登录按钮的自定义显示名称",
	},
}
