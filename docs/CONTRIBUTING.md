# 贡献指南 / Contributing Guide

感谢您对本项目的关注！我们欢迎任何形式的贡献。

Thank you for your interest in this project! We welcome all forms of contributions.

## 🤝 如何贡献 / How to Contribute

### 报告问题 / Report Issues

如果您发现了 bug 或有功能建议：

If you find a bug or have a feature suggestion:

1. 检查 [Issues](https://github.com/starvpn/answer-user-wxcom/issues) 是否已有相同问题
2. 如果没有，创建新的 Issue
3. 详细描述问题或建议
4. 提供复现步骤（如果是 bug）

1. Check if the issue already exists in [Issues](https://github.com/starvpn/answer-user-wxcom/issues)
2. If not, create a new Issue
3. Describe the problem or suggestion in detail
4. Provide reproduction steps (if it's a bug)

### 提交代码 / Submit Code

1. **Fork 本仓库**

   Fork this repository

2. **创建分支**

   Create a branch

```bash
git checkout -b feature/your-feature-name
# 或 / or
git checkout -b fix/your-bug-fix
```

3. **进行开发**

   Develop your changes

4. **编写测试**（如适用）

   Write tests (if applicable)

5. **提交代码**

   Commit your changes

```bash
git add .
git commit -m "feat: add your feature description"
# 或 / or
git commit -m "fix: fix your bug description"
```

6. **推送到您的 Fork**

   Push to your fork

```bash
git push origin feature/your-feature-name
```

7. **创建 Pull Request**

   Create a Pull Request

## 📝 提交信息规范 / Commit Message Convention

我们遵循 [Conventional Commits](https://www.conventionalcommits.org/) 规范：

We follow [Conventional Commits](https://www.conventionalcommits.org/) convention:

- `feat:` 新功能 / New feature
- `fix:` Bug 修复 / Bug fix
- `docs:` 文档更新 / Documentation update
- `style:` 代码格式 / Code formatting
- `refactor:` 重构 / Refactoring
- `test:` 测试 / Testing
- `chore:` 构建/工具 / Build/tooling

**示例 / Examples:**

```
feat: add support for custom redirect URL
fix: resolve state parameter validation issue
docs: update installation guide
refactor: improve error handling
```

## 🔍 代码审查 / Code Review

所有提交都需要经过代码审查。审查者可能会：

All submissions require code review. Reviewers may:

- 请求更改
- 提出建议
- 批准合并

- Request changes
- Provide suggestions
- Approve merge

## ✅ 开发规范 / Development Guidelines

### Go 代码规范 / Go Code Standards

- 遵循 [Effective Go](https://golang.org/doc/effective_go.html)
- 使用 `gofmt` 格式化代码
- 添加必要的注释
- 错误处理要完善

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` for formatting
- Add necessary comments
- Handle errors properly

### 安全性 / Security

- 不要在代码中硬编码敏感信息
- 使用环境变量或配置文件
- 验证所有外部输入
- 使用 HTTPS 进行 API 调用

- Never hardcode sensitive information
- Use environment variables or config files
- Validate all external inputs
- Use HTTPS for API calls

### 测试 / Testing

- 为新功能编写测试
- 确保所有测试通过
- 测试覆盖率保持在合理水平

- Write tests for new features
- Ensure all tests pass
- Maintain reasonable test coverage

## 📚 开发环境设置 / Development Environment Setup

### 前置要求 / Prerequisites

- Go 1.19+
- Git
- Apache Answer 开发环境

### 本地开发 / Local Development

1. Clone 仓库 / Clone repository

```bash
git clone https://github.com/starvpn/answer-user-wxcom.git
cd answer-connector-wechat
```

2. 安装依赖 / Install dependencies

```bash
go mod tidy
```

3. 运行测试 / Run tests

```bash
go test ./...
```

4. 集成到 Answer / Integrate with Answer

参考 README.md 中的安装步骤

Refer to installation steps in README.md

## 🐛 调试技巧 / Debugging Tips

### 启用详细日志 / Enable Verbose Logging

在 Answer 配置中启用 debug 日志：

Enable debug logging in Answer configuration:

```yaml
log:
  level: debug
```

### 测试 OAuth 流程 / Test OAuth Flow

使用工具如 Postman 或 curl 测试 API 端点：

Use tools like Postman or curl to test API endpoints:

```bash
# 获取 access_token
curl -X GET "https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code"

# 获取用户信息
curl -X GET "https://api.weixin.qq.com/sns/userinfo?access_token=TOKEN&openid=OPENID"
```

## 📖 相关资源 / Related Resources

- [Apache Answer 文档](https://answer.apache.org/docs)
- [微信开放平台文档](https://developers.weixin.qq.com/doc/)
- [OAuth 2.0 规范](https://oauth.net/2/)
- [Go 官方文档](https://golang.org/doc/)

## 💬 获取帮助 / Get Help

- 💬 [Discussions](https://github.com/starvpn/answer-user-wxcom/discussions)
- 🐛 [Issues](https://github.com/starvpn/answer-user-wxcom/issues)
- 📧 Email: your-email@example.com

## 📜 行为准则 / Code of Conduct

我们致力于提供友好、安全和包容的环境。请：

We are committed to providing a friendly, safe, and inclusive environment. Please:

- 尊重所有贡献者
- 保持建设性的讨论
- 接受建设性的批评
- 关注项目的最佳利益

- Respect all contributors
- Keep discussions constructive
- Accept constructive criticism
- Focus on what's best for the project

## 🙏 致谢 / Acknowledgments

感谢所有为本项目做出贡献的人！

Thanks to all contributors to this project!

---

再次感谢您的贡献！🎉

Thank you again for your contribution! 🎉

