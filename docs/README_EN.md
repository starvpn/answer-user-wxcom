# Answer WeChat Login Plugin

A WeChat OAuth connector plugin for [Apache Answer](https://github.com/apache/answer) that enables users to sign in with their WeChat account.

## ✨ Features

- 🔐 OAuth 2.0-based WeChat login
- 🌍 Multi-language support (Chinese, English)
- 🎨 Built-in WeChat official logo
- ⚙️ Simple configuration management
- 🔒 Secure state verification (CSRF protection)

## 📋 Prerequisites

Before using this plugin, you need to:

1. Register a developer account on [WeChat Open Platform](https://open.weixin.qq.com/)
2. Create a **Website Application**
3. Submit for review and get approval
4. Obtain the application's **AppID** and **AppSecret**
5. Configure the **Authorization Callback Domain**

### Configure Authorization Callback Domain

In the WeChat Open Platform application management, configure the authorization callback domain as:

```
yourdomain.com
```

**Note:**
- Do not include `http://` or `https://`
- Do not include the path part
- Example: If your Answer is deployed at `https://qa.example.com`, fill in `qa.example.com`

## 🚀 Installation

### Method 1: Install as External Plugin

1. Clone or download this plugin to your project directory:

```bash
git clone https://github.com/starvpn/answer-user-wxcom.git
cd answer-connector-wechat
```

2. Import the plugin in Answer project's `cmd/answer/main.go`:

```go
import (
    answercmd "github.com/apache/answer/cmd"
    // Import WeChat connector plugin
    _ "github.com/apache/answer/plugin/connector-wechat"
)
```

3. Add the plugin to `go.mod` using `go mod edit`:

```bash
go mod edit -replace=github.com/apache/answer/plugin/connector-wechat=/path/to/connector-wechat
```

4. Update dependencies:

```bash
go mod tidy
```

5. Build and run:

```bash
make build
./answer run -C ./answer-data
```

### Method 2: Integrate into Source Code

Copy this plugin code to the `plugin/` directory of Answer source code, then follow steps 2-5 above.

## ⚙️ Configuration

### Configure in Answer Admin Panel

1. Start Answer and log in with an administrator account
2. Go to **Admin** > **Settings** > **Login**
3. Find the **WeChat** connector
4. Fill in the configuration:

| Field | Description | Required |
|-------|-------------|----------|
| AppID | WeChat Open Platform Application ID | ✅ |
| AppSecret | WeChat Open Platform Application Secret | ✅ |
| Display Name | Custom display name for login button | ❌ |

5. Save and enable

### Configuration Example

```json
{
  "app_id": "wx1234567890abcdef",
  "app_secret": "your_app_secret_here",
  "display_name": "Sign in with WeChat"
}
```

## 📖 Usage

After configuration:

1. Users will see the **Sign in with WeChat** button on the login page
2. Clicking the button redirects to WeChat OAuth authorization page
3. Users scan the QR code or use quick login (if WeChat client is logged in)
4. After authorization, users are automatically redirected back to Answer and logged in

## 🔧 Technical Implementation

### OAuth 2.0 Authorization Flow

```
┌─────────┐                                        ┌──────────┐
│  User   │                                        │  WeChat  │
└────┬────┘                                        └─────┬────┘
     │                                                   │
     │  1. Click WeChat Login                            │
     ├──────────────────────────────────────────────────>│
     │                                                   │
     │  2. Show QR Code/Quick Login                      │
     │<──────────────────────────────────────────────────┤
     │                                                   │
     │  3. User Authorizes                               │
     ├──────────────────────────────────────────────────>│
     │                                                   │
     │  4. Return code                                   │
     │<──────────────────────────────────────────────────┤
     │                                                   │
┌────┴────┐                                        ┌─────┴────┐
│ Answer  │  5. Exchange code for access_token     │  WeChat  │
│ Server  ├───────────────────────────────────────>│   API    │
│         │                                        │          │
│         │  6. Return access_token and openid     │          │
│         │<───────────────────────────────────────┤          │
│         │                                        │          │
│         │  7. Get user info with access_token    │          │
│         ├───────────────────────────────────────>│          │
│         │                                        │          │
│         │  8. Return user info (nickname, etc.)  │          │
│         │<───────────────────────────────────────┤          │
└─────────┘                                        └──────────┘
```

## 🔒 Security

### CSRF Protection

The plugin implements state parameter verification to prevent Cross-Site Request Forgery attacks:

1. Generate random state and store in Cookie when initiating authorization
2. Verify that returned state matches Cookie value on callback
3. Reject login request if verification fails

### Sensitive Information Protection

- ⚠️ **AppSecret must be kept secret**, do not commit to version control
- Recommend using environment variables for sensitive configuration
- All API requests use HTTPS encrypted transmission

### UnionID vs OpenID

- **OpenID**: User's unique identifier in current application, different for each app
- **UnionID**: User's unique identifier across all applications under the same open platform account

The plugin uses **UnionID** as the user's unique identifier to ensure user consistency across applications.

## 📝 Notes

1. **Email Field**: WeChat OAuth doesn't provide user email, this field is empty
2. **Username Generation**: Auto-generated in format `wechat_<first8chars_of_openid>`
3. **Avatar**: Uses avatar URL provided by WeChat
4. **First Login**: Account is automatically created in Answer on first WeChat login

## 🐛 Troubleshooting

### Common Issues

**Q: "This link is inaccessible" error**

A: Please check:
- Is AppID correct
- Does redirect_uri domain match the authorization callback domain configured in WeChat Open Platform
- Is scope set to `snsapi_login`

**Q: Failed to get user information**

A: Please check:
- Is AppSecret correct
- Has the application passed review
- Is WeChat login functionality enabled

**Q: Incorrect callback address**

A: Please ensure:
- Answer domain is configured correctly
- Firewall/reverse proxy allows WeChat server access
- Callback URL format is correct

## 📄 License

Apache License 2.0

## 🤝 Contributing

Issues and Pull Requests are welcome!

## 📚 References

- [Apache Answer Plugin Development Guide](https://answer.apache.org/docs/development/plugins/)
- [WeChat Open Platform Documentation](https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Wechat_Login.html)
- [OAuth 2.0 Specification](https://oauth.net/2/)

## 💬 Support

For questions:
- Submit an [Issue](https://github.com/starvpn/answer-user-wxcom/issues)
- Visit [Answer Community](https://answer.apache.org/community)

