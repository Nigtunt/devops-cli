# 贡献指南

欢迎为 DevOps CLI 贡献代码！

## 开发环境设置

```bash
# 克隆仓库
git clone https://github.com/Nigtunt/devops-cli.git
cd devops-cli

# 安装依赖
make deps

# 编译
make build

# 运行测试
make test
```

## 添加新命令

1. 在 `cmd/` 目录下创建新的资源目录，例如 `cmd/issue/`
2. 创建 `issue.go` 文件，参考 `cmd/task/task.go` 的结构
3. 在 `main.go` 中注册新命令

## 代码风格

- 使用 `go fmt` 格式化代码
- 遵循 Go 标准命名规范
- 命令使用小写，单词间无分隔符

## 提交 PR

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 发布新版本

```bash
# 打 tag
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0

# 编译所有平台
make package
```

## 问题反馈

遇到问题？请提 [Issue](https://github.com/Nigtunt/devops-cli/issues)
