# 项目结构

```
answer-user-wxcom/
│
├── 核心代码
│   ├── wechat.go                # 主实现（OAuth 流程）
│   ├── i18n.go                  # 国际化常量
│   └── i18n/                    # 翻译文件
│       ├── translation.yaml     # 英文
│       └── zh_CN.yaml           # 中文
│
├── 配置
│   ├── go.mod                   # Go 模块定义
│   ├── info.yaml                # 插件元信息
│   └── config.example.yaml      # 配置示例
│
├── 文档
│   ├── README.md                # 主文档 ⭐
│   └── docs/                    # 详细文档
│       ├── QUICKSTART.md        # 快速开始
│       ├── README_EN.md         # 英文文档
│       ├── CONTRIBUTING.md      # 贡献指南
│       └── CHANGELOG.md         # 更新日志
│
├── 工具
│   ├── install.sh               # 安装脚本
│   └── .gitignore               # Git 配置
│
└── LICENSE                      # Apache 2.0 许可证
```

## 文件说明

### 核心文件
- `wechat.go` (~260 行) - 插件核心实现，包含 OAuth 2.0 流程
- `i18n.go` (~60 行) - 国际化翻译键定义
- `go.mod` - Go 模块定义，module path: `github.com/starvpn/answer-user-wxcom`

### 配置文件
- `info.yaml` - 插件元信息（名称、版本、作者等）
- `config.example.yaml` - 配置示例文件

### 文档文件
- `README.md` - 项目主页，包含快速开始和基本用法
- `docs/QUICKSTART.md` - 详细的安装配置步骤
- `docs/README_EN.md` - 英文版文档
- `docs/CONTRIBUTING.md` - 如何参与贡献
- `docs/CHANGELOG.md` - 版本更新历史

## 快速导航

- 想要快速安装？→ [README.md](../README.md#快速安装)
- 想要详细步骤？→ [QUICKSTART.md](./QUICKSTART.md)
- 想要参与开发？→ [CONTRIBUTING.md](./CONTRIBUTING.md)
- 英文文档？→ [README_EN.md](./README_EN.md)

