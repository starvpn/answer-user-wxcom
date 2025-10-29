#!/bin/bash

# Answer 微信登录插件安装脚本
# WeChat Login Plugin Installation Script for Answer
#
# 使用方法 / Usage:
#   方式 1 / Method 1 (推荐 / Recommended): answer build --with github.com/starvpn/answer-user-wxcom
#   方式 2 / Method 2: ./install.sh /path/to/answer

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 打印带颜色的消息
print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 检查必要的命令
check_requirements() {
    print_info "检查系统要求..."
    
    if ! command -v go &> /dev/null; then
        print_error "未找到 Go 编译器，请先安装 Go 1.19+"
        exit 1
    fi
    
    GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
    print_success "检测到 Go 版本: $GO_VERSION"
}

# 获取 Answer 项目路径
get_answer_path() {
    if [ -z "$1" ]; then
        print_error "请提供 Answer 项目路径"
        echo ""
        echo "使用方法:"
        echo "  ./install.sh /path/to/answer"
        echo ""
        exit 1
    fi
    
    ANSWER_PATH="$1"
    
    if [ ! -d "$ANSWER_PATH" ]; then
        print_error "Answer 项目路径不存在: $ANSWER_PATH"
        exit 1
    fi
    
    if [ ! -f "$ANSWER_PATH/go.mod" ]; then
        print_error "指定路径不是有效的 Go 项目: $ANSWER_PATH"
        exit 1
    fi
    
    print_success "找到 Answer 项目: $ANSWER_PATH"
}

# 安装插件
install_plugin() {
    print_info "开始安装微信登录插件..."
    
    # 创建插件目录
    PLUGIN_DIR="$ANSWER_PATH/plugin/connector-wechat"
    mkdir -p "$PLUGIN_DIR"
    
    # 复制插件文件
    print_info "复制插件文件..."
    cp -r connector-wechat/* "$PLUGIN_DIR/"
    print_success "插件文件已复制"
    
    # 更新 go.mod
    print_info "更新 go.mod..."
    cd "$ANSWER_PATH"
    go mod edit -replace=github.com/apache/answer/plugin/connector-wechat=./plugin/connector-wechat
    print_success "go.mod 已更新"
    
    # 检查 main.go
    MAIN_GO="$ANSWER_PATH/cmd/answer/main.go"
    if [ ! -f "$MAIN_GO" ]; then
        print_error "未找到 main.go: $MAIN_GO"
        exit 1
    fi
    
    # 检查是否已经导入
    if grep -q "connector-wechat" "$MAIN_GO"; then
        print_warning "检测到已经导入插件，跳过..."
    else
        print_info "更新 main.go..."
        # 备份原文件
        cp "$MAIN_GO" "$MAIN_GO.backup"
        
        # 添加导入语句 (需要手动完成)
        print_warning "请手动在 $MAIN_GO 中添加以下导入语句:"
        echo ""
        echo "  import ("
        echo "      answercmd \"github.com/apache/answer/cmd\""
        echo "      // 导入微信登录插件"
        echo "      _ \"github.com/apache/answer/plugin/connector-wechat\""
        echo "  )"
        echo ""
        print_info "已备份原文件到: $MAIN_GO.backup"
    fi
    
    # 更新依赖
    print_info "更新依赖..."
    go mod tidy
    print_success "依赖已更新"
}

# 构建项目
build_project() {
    print_info "构建 Answer..."
    cd "$ANSWER_PATH"
    
    if [ -f "Makefile" ]; then
        make build
    else
        go build -o answer cmd/answer/main.go
    fi
    
    if [ $? -eq 0 ]; then
        print_success "构建成功！"
    else
        print_error "构建失败，请检查错误信息"
        exit 1
    fi
}

# 显示下一步操作
show_next_steps() {
    echo ""
    print_success "=========================================="
    print_success "微信登录插件安装完成！"
    print_success "=========================================="
    echo ""
    print_info "下一步操作:"
    echo ""
    echo "1. 确保已在 cmd/answer/main.go 中导入插件"
    echo "   import _ \"github.com/apache/answer/plugin/connector-wechat\""
    echo ""
    echo "2. 启动 Answer:"
    echo "   cd $ANSWER_PATH"
    echo "   ./answer run -C ./answer-data"
    echo ""
    echo "3. 登录管理后台配置插件:"
    echo "   - 访问: https://yourdomain.com/admin"
    echo "   - 导航到: 设置 -> 登录 -> 微信"
    echo "   - 填写 AppID 和 AppSecret"
    echo "   - 保存并启用"
    echo ""
    echo "4. 查看完整文档:"
    echo "   - 快速开始: connector-wechat/QUICKSTART.md"
    echo "   - 完整文档: connector-wechat/README.md"
    echo ""
    print_warning "注意: 请确保已在微信开放平台配置授权回调域！"
    echo ""
}

# 主函数
main() {
    echo ""
    echo "=========================================="
    echo "  Answer 微信登录插件安装脚本"
    echo "  WeChat Login Plugin Installer"
    echo "=========================================="
    echo ""
    
    check_requirements
    get_answer_path "$1"
    install_plugin
    
    # 询问是否立即构建
    echo ""
    read -p "是否立即构建项目? (y/n): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        build_project
    else
        print_info "跳过构建，您可以稍后手动构建"
    fi
    
    show_next_steps
}

# 显示帮助信息
show_help() {
    echo "使用方法:"
    echo "  ./install.sh <answer-path>"
    echo ""
    echo "参数:"
    echo "  answer-path    Answer 项目的根目录路径"
    echo ""
    echo "示例:"
    echo "  ./install.sh /home/user/answer"
    echo "  ./install.sh ../answer"
    echo ""
    echo "选项:"
    echo "  -h, --help     显示此帮助信息"
    echo ""
}

# 处理命令行参数
if [ "$1" = "-h" ] || [ "$1" = "--help" ]; then
    show_help
    exit 0
fi

# 运行主函数
main "$@"

