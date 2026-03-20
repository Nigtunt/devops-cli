# DevOps CLI

基于 Cobra 框架开发的平台 API 命令行工具。

## 安装

### 方式一：从源码编译

```bash
make build
sudo make install
```

### 方式二：下载预编译二进制

从 [Releases](https://github.com/your-org/devops-cli/releases) 下载对应平台的压缩包：

- **Linux**: `yx-linux-amd64.tar.gz` / `yx-linux-arm64.tar.gz` / `yx-linux-arm.tar.gz`
- **Windows**: `yx-windows-amd64.zip` / `yx-windows-arm64.zip`
- **macOS**: `yx-darwin-amd64.tar.gz` / `yx-darwin-arm64.tar.gz`

解压后放到 PATH 即可：
```bash
tar -xzf yx-linux-amd64.tar.gz
sudo mv yx-linux-amd64 /usr/local/bin/yx
```

### 方式三：多平台编译

```bash
# 编译所有平台
make build-all

# 编译并打包
make package
```

## 配置

### 方式一：使用 token 登录（推荐）

```bash
# 登录后 token 会自动保存到 ~/.devops-cli.yaml
yx auth login --token <your_token>

# 查看登录状态
yx auth status

# 登出
yx auth logout
```

### 方式二：手动配置

在 `~/.devops-cli.yaml` 创建配置文件:

```yaml
api_token: your_token_here
base_url: https://api.example.com
```

### 方式三：环境变量

```bash
export DEVOPS_TOKEN=<your_token>
export DEVOPS_API_URL=https://api.example.com
```

## 使用

### 认证

```bash
# 登录（token 会加密存储）
yx auth login --token <your_token>

# 查看状态
yx auth status

# 登出
yx auth logout
```

### Debug 模式

```bash
# 启用 debug 模式（显示详细请求/响应，敏感信息自动脱敏）
yx userstory create -t "标题" -d "描述" --debug

# 或简写
yx userstory create -t "标题" -d "描述" -D
```

### 输出格式化

```bash
# 表格输出（默认）
yx userstory list

# JSON 输出
yx userstory list --format json

# YAML 输出
yx userstory list --format yaml
```

### 用户故事

```bash
# 创建用户故事
yx userstory create --title "用户登录功能" --description "作为用户，我希望能够登录系统"

# 简写
yx userstory create -t "标题" -d "描述"
```

## 目录结构

```
devops-cli/
├── cmd/
│   ├── root.go                  # 根命令
│   ├── userstory/
│   │   └── userstory.go         # userstory create/list/view
│   ├── systemchange/
│   │   └── systemchange.go      # systemchange create/list/view/edit
│   └── task/
│       └── task.go              # task create/list/view
├── internal/
│   ├── config/                  # 配置管理
│   ├── api/                     # API 客户端 (待实现)
│   └── pkg/                     # 公共包
├── main.go                      # 入口
├── go.mod
└── README.md
```

## 扩展新命令

在 `cmd/` 目录下创建新的命令文件，参考 `userstory.go` 的结构。
