# 更新日志

所有重要的项目变更都会记录在此文件中。

## [1.0.3] - 2025-10-30

### 新增 ✨
- 完整的微信 OAuth 2.0 登录实现
- 支持微信开放平台网站应用
- 多语言支持（中文和英文）
- 自动用户创建和信息同步
- 微信官方品牌 Logo
- 完善的配置管理界面

### 安全 🔒
- 实现 CSRF 防护（State 参数验证）
- UnionID 空值容错处理
- 数组越界保护
- 微信 API 错误码检查
- AppSecret 使用密码输入类型
- HTTPS 强制加密通信
- HttpOnly Cookie 安全存储

### 修复 🐛
- 修复 UnionID 为空时登录失败的问题
- 修复 OpenID 长度不足导致的 panic
- 修复 getUserInfo 未检查错误码的问题
- 修复 AppSecret 明文显示的问题

### 优化 ⚡
- 添加语言参数（lang=zh_CN）
- 优化用户名生成逻辑
- 改进错误提示信息
- 完善日志记录

### 文档 📖
- 中英文双语 README
- 详细的快速开始指南
- 贡献指南
- 安全检查清单
- 项目结构说明

### 技术细节 🔧
- Go 1.19+
- Answer 1.3.0+
- OAuth 2.0 授权码模式
- RESTful API 调用

---

## 计划功能

### [1.1.0] - 计划中
- [ ] 支持企业微信扫码登录
- [ ] 内嵌式二维码 UI
- [ ] 更多配置选项
- [ ] 登录统计功能

### [2.0.0] - 未来
- [ ] 微信小程序登录支持
- [ ] 账号绑定/解绑功能
- [ ] 多账号绑定
- [ ] 社交图谱集成

---

## 版本说明

版本号遵循 [语义化版本](https://semver.org/lang/zh-CN/) 规范：

- **MAJOR（主版本）**：不兼容的 API 变更
- **MINOR（次版本）**：向后兼容的功能新增
- **PATCH（修订版）**：向后兼容的问题修复

---

## 链接

- [项目主页](https://github.com/starvpn/answer-user-wxcom)
- [问题反馈](https://github.com/starvpn/answer-user-wxcom/issues)
- [拉取请求](https://github.com/starvpn/answer-user-wxcom/pulls)
