# Answer 微信登录插件

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Version](https://img.shields.io/badge/Go-1.19%2B-blue.svg)](https://golang.org/dl/)
[![Answer Version](https://img.shields.io/badge/Answer-1.3.0%2B-green.svg)](https://github.com/apache/answer)

为 [Apache Answer](https://github.com/apache/answer) 开发的微信 OAuth 登录插件，支持微信扫码登录。

---

## ✨ 特性

- 🔐 OAuth 2.0 标准登录
- 🌍 多语言支持（中文/英文）
- 🎨 微信官方品牌 Logo
- ⚙️ 简单配置（AppID + AppSecret）
- 🔒 CSRF 防护和安全验证
- 📱 支持扫码和 PC 快速登录

---

## 🚀 快速安装

### 一条命令安装（推荐）

```bash
answer build \
    --with github.com/starvpn/answer-user-wxcom \
    --output ./new_answer

# 替换原有的 answer
sudo mv ./new_answer /usr/bin/answer

# 启动 Answer
answer run -C ./answer-data
```

### Docker 部署

```dockerfile
FROM answerdev/answer:latest

RUN answer build \
    --with github.com/starvpn/answer-user-wxcom \
    --output /usr/bin/answer
```

```bash
docker build -t answer-wechat .
docker run -d -p 80:80 -v ./data:/data answer-wechat
```

### 安装多个插件

```bash
answer build \
    --with github.com/starvpn/answer-user-wxcom \
    --with github.com/apache/answer-plugins/storage-s3 \
    --with github.com/apache/answer-plugins/search-elasticsearch \
    --output ./new_answer
```

---

## ⚙️ 配置

### 1. 微信开放平台准备

1. 注册[微信开放平台](https://open.weixin.qq.com/)开发者账号
2. 创建**网站应用**并通过审核
3. 获取 **AppID** 和 **AppSecret**
4. 配置**授权回调域**（如：`yourdomain.com`，不含协议和路径）

### 2. Answer 后台配置

1. 登录 Answer 管理后台
2. 导航到：**设置** → **登录** → **微信**
3. 填写配置并保存：

| 配置项 | 必填 | 说明 |
|-------|------|------|
| AppID | ✅ | 微信开放平台应用 ID |
| AppSecret | ✅ | 微信开放平台应用密钥 |
| Display Name | ❌ | 自定义按钮显示名称 |

4. 启用插件

### 3. 测试登录

访问 Answer 登录页面，应该能看到"使用微信登录"按钮。点击后扫码即可登录。

---

## 📖 文档

- 📘 [快速开始指南](./docs/QUICKSTART.md) - 详细的安装和配置步骤
- 🌍 [English Documentation](./docs/README_EN.md) - English version
- 🤝 [贡献指南](./docs/CONTRIBUTING.md) - 如何参与贡献
- 📝 [更新日志](./docs/CHANGELOG.md) - 版本历史

---

## 🔧 开发

### 本地开发

```bash
# 克隆项目
git clone https://github.com/starvpn/answer-user-wxcom.git

# 链接到 Answer 项目
cd /path/to/answer
go mod edit -replace=github.com/starvpn/answer-user-wxcom=/path/to/answer-user-wxcom

# 在 cmd/answer/main.go 中导入
import _ "github.com/starvpn/answer-user-wxcom"

# 构建
go mod tidy
make build
```

### 项目结构

```
answer-user-wxcom/
├── wechat.go              # 核心实现
├── i18n.go                # 国际化
├── go.mod                 # Go 模块
├── i18n/                  # 翻译文件
├── docs/                  # 文档目录
└── config.example.yaml    # 配置示例
```

---

## 🐛 故障排查

### 常见问题

**Q: "该链接无法访问"**  
A: 检查 AppID 是否正确，授权回调域是否配置正确（不要包含 `https://` 和路径）

**Q: 获取用户信息失败**  
A: 检查 AppSecret 是否正确，应用是否已通过审核

**Q: 回调地址错误**  
A: 确保 Answer 的域名配置正确，且防火墙允许微信服务器访问

详细故障排查请查看：[快速开始指南](./docs/QUICKSTART.md#故障排查)

---

## 🔒 安全说明

- ⚠️ **AppSecret 必须保密**，不要提交到版本控制系统
- ✅ 生产环境必须使用 HTTPS
- ✅ 已实现 CSRF 防护（State 参数验证）
- ✅ 使用安全的 Cookie 处理

---

## 🤝 贡献

欢迎贡献代码、报告问题、提出建议！

查看 [贡献指南](./docs/CONTRIBUTING.md) 了解详情。

---

## 📄 许可证

[Apache License 2.0](./LICENSE)

---

## 🙏 致谢

- [Apache Answer](https://github.com/apache/answer) - 优秀的问答平台
- [WeChat Open Platform](https://open.weixin.qq.com/) - 提供 OAuth API

---

## 📞 支持

- 🐛 [提交 Issue](https://github.com/starvpn/answer-user-wxcom/issues)
- 💬 [GitHub Discussions](https://github.com/starvpn/answer-user-wxcom/discussions)
- 🏠 [Answer 官网](https://answer.apache.org)
- 📖 [微信开放平台文档](https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Wechat_Login.html)

---

<div align="center">

**Made with ❤️ for Answer Community**

[⭐ Star](https://github.com/starvpn/answer-user-wxcom) · [📖 Docs](./docs/QUICKSTART.md) · [🐛 Issues](https://github.com/starvpn/answer-user-wxcom/issues)

</div>
