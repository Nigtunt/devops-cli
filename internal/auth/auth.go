package auth

import (
	"fmt"
	"os"
	"path/filepath"

	"devops-cli/internal/api"
	"devops-cli/internal/config"
	"gopkg.in/yaml.v3"
)

// Login 登录并保存 token
func Login(token string) error {
	// 创建客户端验证 token
	client := api.NewClient()
	client.SetToken(token)

	fmt.Println("🔐 正在验证 token...")
	user, err := client.ValidateToken()
	if err != nil {
		return fmt.Errorf("token 验证失败：%w", err)
	}

	// 保存配置
	cfgFile := getConfigPath()
	cfg := &config.Config{
		APIToken: token,
		BaseURL:  config.Get().BaseURL,
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("序列化配置失败：%w", err)
	}

	if err := os.WriteFile(cfgFile, data, 0600); err != nil {
		return fmt.Errorf("保存配置失败：%w", err)
	}

	fmt.Printf("✅ 登录成功！欢迎, %s (%s)\n", user.Username, user.Email)
	fmt.Printf("📁 配置已保存到：%s\n", cfgFile)
	return nil
}

// Logout 登出
func Logout() error {
	cfgFile := getConfigPath()
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		return fmt.Errorf("未登录")
	}

	// 清空 token
	cfg := &config.Config{
		BaseURL: config.Get().BaseURL,
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("序列化配置失败：%w", err)
	}

	if err := os.WriteFile(cfgFile, data, 0600); err != nil {
		return fmt.Errorf("清除 token 失败：%w", err)
	}

	fmt.Println("✅ 已登出")
	return nil
}

// Status 显示认证状态
func Status() error {
	cfg := config.Get()
	if cfg.APIToken == "" {
		fmt.Println("❌ 未登录")
		fmt.Println("\n💡 使用以下命令登录:")
		fmt.Println("   yx auth login --token <your_token>")
		return nil
	}

	// 验证 token
	client := api.NewClient()
	user, err := client.ValidateToken()
	if err != nil {
		fmt.Println("❌ Token 已失效")
		fmt.Println("\n💡 使用以下命令重新登录:")
		fmt.Println("   yx auth login --token <new_token>")
		return nil
	}

	fmt.Println("✅ 已登录")
	fmt.Printf("👤 用户：%s (%s)\n", user.Username, user.Email)
	fmt.Printf("🔧 API 地址：%s\n", cfg.BaseURL)
	return nil
}

// getConfigPath 获取配置文件路径
func getConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "无法获取用户主目录:", err)
		os.Exit(1)
	}
	return filepath.Join(home, ".devops-cli.yaml")
}
