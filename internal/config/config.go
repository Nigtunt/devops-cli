package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config 配置结构
type Config struct {
	APIKey    string `yaml:"api_key"`
	APISecret string `yaml:"api_secret"`
	APIToken  string `yaml:"api_token"`
	BaseURL   string `yaml:"base_url"`
	Debug     bool   `yaml:"debug"`
}

var cfg *Config

// Init 初始化配置
func Init(cfgFile string) {
	if cfgFile == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, "无法获取用户主目录:", err)
			os.Exit(1)
		}
		cfgFile = filepath.Join(home, ".devops-cli.yaml")
	}

	cfg = &Config{
		BaseURL: "https://api.example.com",
	}

	if _, err := os.Stat(cfgFile); err == nil {
		data, err := os.ReadFile(cfgFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "读取配置文件失败:", err)
			os.Exit(1)
		}
		if err := yaml.Unmarshal(data, cfg); err != nil {
			fmt.Fprintln(os.Stderr, "解析配置文件失败:", err)
			os.Exit(1)
		}
	}
}

// Get 获取配置
func Get() *Config {
	return cfg
}
