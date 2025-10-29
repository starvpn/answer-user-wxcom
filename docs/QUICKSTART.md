# 快速开始 / Quick Start Guide

本指南将帮助您在 5 分钟内完成微信登录插件的配置。

This guide will help you configure the WeChat login plugin in 5 minutes.

---

## 🚀 5 分钟快速配置 / 5-Minute Quick Setup

### 步骤 1: 获取微信开放平台凭证 / Step 1: Get WeChat Credentials

1. 访问 [微信开放平台](https://open.weixin.qq.com/)
2. 注册并登录开发者账号
3. 创建"网站应用"
4. 获取 **AppID** 和 **AppSecret**

**重要：** 记录以下信息

**Important:** Record the following information

```
AppID: wx__________________ (18位)
AppSecret: ________________________________ (32位)
```

### 步骤 2: 配置授权回调域 / Step 2: Configure Callback Domain

在微信开放平台的应用管理中：

In WeChat Open Platform application management:

```
授权回调域 / Callback Domain: yourdomain.com
```

**示例 / Example:**
- ✅ 正确: `qa.example.com`
- ❌ 错误: `https://qa.example.com`
- ❌ 错误: `qa.example.com/callback`

### 步骤 3: 安装插件 / Step 3: Install Plugin

#### 方式 A: 使用 answer build 命令（最简单）⭐

**Approach A: Using answer build Command (Easiest)**

这是最简单和推荐的方式！

This is the easiest and recommended way!

```bash
# 1. 确保插件已发布到 GitHub
# Ensure plugin is published on GitHub
# 例如 / For example: github.com/starvpn/answer-user-wxcom

# 2. 使用 answer build 构建
# Build with answer build command
answer build \
    --with github.com/starvpn/answer-user-wxcom \
    --output ./new_answer

# 3. 替换原有的 answer
# Replace the original answer binary
mv ./new_answer /usr/bin/answer
# 或者如果在当前目录 / Or if in current directory:
# mv ./new_answer ./answer
```

**一键安装完成！无需修改代码！**

**Done in one command! No code modification needed!**

#### 方式 B: Docker 部署（推荐生产环境）

**Approach B: Docker Deployment (Recommended for Production)**

创建 `Dockerfile`:

Create `Dockerfile`:

```dockerfile
FROM answerdev/answer:latest

RUN answer build \
    --with github.com/starvpn/answer-user-wxcom \
    --output /usr/bin/answer
```

构建并运行:

Build and run:

```bash
docker build -t answer-with-wechat .
docker run -d -p 80:80 answer-with-wechat
```

#### 方式 C: 使用 go mod replace（开发环境）

**Approach C: Using go mod replace (Development)**

适用于本地开发和测试。

For local development and testing.

```bash
# 1. 在 Answer 项目根目录
cd /path/to/answer

# 2. 克隆插件
git clone https://github.com/starvpn/answer-user-wxcom.git plugins/connector-wechat

# 3. 添加 replace 指令
go mod edit -replace=github.com/starvpn/answer-user-wxcom=./plugins/connector-wechat

# 4. 在 cmd/answer/main.go 中添加导入
# Add import in cmd/answer/main.go
```

编辑 `cmd/answer/main.go`:

Edit `cmd/answer/main.go`:

```go
package main

import (
    answercmd "github.com/apache/answer/cmd"
    
    // 导入微信登录插件
    // Import WeChat connector plugin
    _ "github.com/starvpn/answer-user-wxcom"
)

func main() {
    answercmd.Main()
}
```

```bash
# 5. 更新依赖并构建
go mod tidy
make build
```

### 步骤 4: 启动 Answer / Step 4: Start Answer

```bash
./answer run -C ./answer-data
```

或使用 Docker:

Or using Docker:

```bash
docker-compose up -d
```

### 步骤 5: 配置插件 / Step 5: Configure Plugin

1. 浏览器访问 Answer 管理后台
   
   Visit Answer admin panel in browser

```
https://yourdomain.com/admin
```

2. 导航到: **设置** → **登录** → **微信**

   Navigate to: **Settings** → **Login** → **WeChat**

3. 填写配置:

   Fill in configuration:

| 字段 / Field | 值 / Value | 说明 / Description |
|-------------|-----------|-------------------|
| AppID | `wx__________` | 从步骤 1 获取 / From Step 1 |
| AppSecret | `____________` | 从步骤 1 获取 / From Step 1 |
| Display Name | `微信登录` | 可选 / Optional |

4. 点击 **保存** 并 **启用**

   Click **Save** and **Enable**

### 步骤 6: 测试登录 / Step 6: Test Login

1. 退出当前账号（如已登录）

   Logout current account (if logged in)

2. 访问登录页面

   Visit login page

```
https://yourdomain.com/users/login
```

3. 应该看到 **使用微信登录** 按钮

   Should see **Sign in with WeChat** button

4. 点击按钮测试登录流程

   Click button to test login flow

---

## ✅ 验证清单 / Verification Checklist

配置完成后，检查以下项目：

After configuration, check the following:

- [ ] 插件在管理后台显示
- [ ] AppID 和 AppSecret 已正确填写
- [ ] 插件已启用（开关为绿色）
- [ ] 登录页面显示微信登录按钮
- [ ] 点击按钮可以跳转到微信授权页面
- [ ] 扫码后可以成功登录

- [ ] Plugin shows in admin panel
- [ ] AppID and AppSecret filled correctly
- [ ] Plugin enabled (switch is green)
- [ ] WeChat login button shows on login page
- [ ] Clicking button redirects to WeChat auth page
- [ ] Can login successfully after scanning QR code

---

## 🐛 常见问题快速修复 / Quick Fixes

### ❌ 插件不显示

**Plugin Not Showing**

```bash
# 检查插件是否正确导入
grep "connector-wechat" cmd/answer/main.go

# 重新构建
make build

# 重启服务
./answer restart
```

### ❌ "该链接无法访问"

**"This Link is Inaccessible"**

检查微信开放平台配置：

Check WeChat Open Platform configuration:

1. AppID 是否正确
2. 授权回调域是否匹配
3. 应用是否已审核通过

1. Is AppID correct
2. Does callback domain match
3. Has application been approved

### ❌ 获取用户信息失败

**Failed to Get User Info**

```bash
# 检查日志
tail -f answer-data/logs/answer.log

# 验证 AppSecret 是否正确
# Verify AppSecret is correct
```

### ❌ 回调地址错误

**Wrong Callback URL**

确保 Answer 的域名配置正确：

Ensure Answer domain is configured correctly:

```yaml
# answer-data/config.yaml
server:
  http:
    addr: 0.0.0.0:80
  # 确保这个域名正确
  # Ensure this domain is correct
  hostname: "https://yourdomain.com"
```

---

## 🎯 下一步 / Next Steps

配置成功后，您可以：

After successful configuration, you can:

1. 📝 自定义登录按钮显示名称
2. 🎨 调整登录页面样式
3. 🔧 查看 [完整文档](../README.md)
4. 💡 探索 [高级配置](./UI_CUSTOMIZATION.md)

1. 📝 Customize login button display name
2. 🎨 Adjust login page styling
3. 🔧 Read [Full Documentation](../README.md)
4. 💡 Explore [Advanced Configuration](./UI_CUSTOMIZATION.md)

---

## 📞 需要帮助？ / Need Help?

- 📖 查看 [完整 README](../README.md)
- 🐛 提交 [Issue](https://github.com/starvpn/answer-user-wxcom/issues)
- 💬 访问 [Answer 社区](https://answer.apache.org/community)
- 📧 发送邮件: support@example.com

---

## 🎉 恭喜！

您已经成功配置了微信登录插件！

You have successfully configured the WeChat login plugin!

如果遇到问题，请参考 [故障排查指南](../README.md#-故障排查)。

If you encounter issues, please refer to [Troubleshooting Guide](../README.md#-troubleshooting).

