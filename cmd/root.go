package cmd

import (
	"fmt"
	"os"

	"devops-cli/internal/config"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	debug   bool
	rootCmd = &cobra.Command{
		Use:   "yx",
		Short: "DevOps CLI - 平台接口调用工具",
		Long: `DevOps CLI 是一个用于封装平台 API 调用的命令行工具。
支持用户故事、系统变更、任务等多种资源的管理操作。`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if debug {
				config.Get().Debug = true
			}
		},
	}
)

// RootCmd 导出根命令供 main.go 使用
func RootCmd() *cobra.Command {
	return rootCmd
}

// Execute 执行根命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "配置文件路径 (默认 $HOME/.devops-cli.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "D", false, "启用 debug 模式（显示详细请求/响应）")
}

// initConfig 初始化配置
func initConfig() {
	config.Init(cfgFile)
}
