.PHONY: build test clean run help deps fmt lint build-all clean-dist

# 项目名称
BINARY_NAME=yx
GO=go
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

# 默认目标
all: build

# 安装依赖
deps:
	export GOPROXY=https://goproxy.cn,direct
	$(GO) mod download
	$(GO) mod tidy

# 编译（当前平台）
build: deps
	$(GO) build -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)" -o $(BINARY_NAME) main.go
	@echo "✅ 编译完成：./$(BINARY_NAME)"

# 多平台多架构编译
build-all: clean-dist deps
	@echo "🔨 开始多平台编译..."
	@echo ""
	# Linux amd64
	GOOS=linux GOARCH=amd64 $(GO) build -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)" -o dist/$(BINARY_NAME)-linux-amd64 main.go
	@echo "✅ linux-amd64"
	# Linux arm64
	GOOS=linux GOARCH=arm64 $(GO) build -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)" -o dist/$(BINARY_NAME)-linux-arm64 main.go
	@echo "✅ linux-arm64"
	# Linux arm (32-bit)
	GOOS=linux GOARCH=arm GOARM=7 $(GO) build -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)" -o dist/$(BINARY_NAME)-linux-arm main.go
	@echo "✅ linux-arm"
	# Windows amd64
	GOOS=windows GOARCH=amd64 $(GO) build -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)" -o dist/$(BINARY_NAME)-windows-amd64.exe main.go
	@echo "✅ windows-amd64"
	# Windows arm64
	GOOS=windows GOARCH=arm64 $(GO) build -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)" -o dist/$(BINARY_NAME)-windows-arm64.exe main.go
	@echo "✅ windows-arm64"
	# macOS amd64
	GOOS=darwin GOARCH=amd64 $(GO) build -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)" -o dist/$(BINARY_NAME)-darwin-amd64 main.go
	@echo "✅ darwin-amd64"
	# macOS arm64 (M1/M2)
	GOOS=darwin GOARCH=arm64 $(GO) build -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)" -o dist/$(BINARY_NAME)-darwin-arm64 main.go
	@echo "✅ darwin-arm64"
	@echo ""
	@echo "📦 所有平台编译完成！"
	@echo "📁 输出目录：dist/"
	@ls -lh dist/

# 运行测试
test:
	$(GO) test -v ./...

# 格式化代码
fmt:
	$(GO) fmt ./...

# 代码检查
lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "⚠️  golangci-lint 未安装，跳过"; \
	fi

# 清理（仅编译产物）
clean:
	rm -f $(BINARY_NAME)
	@echo "✅ 清理完成"

# 清理所有（包括 dist）
clean-dist:
	rm -rf dist/ $(BINARY_NAME)
	@echo "✅ 清理完成"

# 运行 CLI
run: build
	./$(BINARY_NAME) --help

# 安装到系统路径
install: build
	sudo mv $(BINARY_NAME) /usr/local/bin/
	@echo "✅ 已安装到 /usr/local/bin/$(BINARY_NAME)"

# 卸载
uninstall:
	sudo rm -f /usr/local/bin/$(BINARY_NAME)
	@echo "✅ 已卸载"

# 打包 tarball
package: build-all
	@echo "📦 打包发布文件..."
	cd dist && for f in yx-*; do \
		case "$$f" in \
			*.exe) \
				zip -q $${f%.exe}.zip $$f && rm $$f ;; \
			*) \
				tar -czf $$f.tar.gz $$f && rm $$f ;; \
		esac; \
	done
	@echo "✅ 打包完成"
	@ls -lh dist/

# 帮助
help:
	@echo "DevOps CLI - Makefile 命令:"
	@echo ""
	@echo "  make build      - 编译当前平台"
	@echo "  make build-all  - 多平台编译 (Linux/Windows/macOS + amd64/arm64)"
	@echo "  make package    - 多平台编译并打包成 tar.gz/zip"
	@echo "  make test       - 运行测试"
	@echo "  make clean      - 清理编译产物"
	@echo "  make clean-dist - 清理所有 (包括 dist/)"
	@echo "  make run        - 编译并运行 --help"
	@echo "  make deps       - 安装依赖"
	@echo "  make fmt        - 格式化代码"
	@echo "  make lint       - 代码检查"
	@echo "  make install    - 安装到系统路径"
	@echo "  make uninstall  - 从系统路径卸载"
	@echo "  make help       - 显示帮助"
