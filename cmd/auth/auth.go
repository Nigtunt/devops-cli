package auth

import (
	"fmt"
	"os"

	"devops-cli/internal/auth"

	"github.com/spf13/cobra"
)

// NewAuthCmd 创建 auth 父命令
func NewAuthCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "认证管理",
		Long:  "认证相关操作，包括登录、登出、查看状态等",
	}

	cmd.AddCommand(loginCmd())
	cmd.AddCommand(logoutCmd())
	cmd.AddCommand(statusCmd())

	return cmd
}

func loginCmd() *cobra.Command {
	var token string

	cmd := &cobra.Command{
		Use:   "login",
		Short: "登录",
		Long:  "使用 token 登录到平台",
		Run: func(cmd *cobra.Command, args []string) {
			// 如果没传 token，尝试从环境变量读取
			if token == "" {
				token = os.Getenv("DEVOPS_TOKEN")
			}

			if token == "" {
				fmt.Println("❌ 请提供 token")
				fmt.Println("\n使用方式:")
				fmt.Println("  1. yx auth login --token <your_token>")
				fmt.Println("  2. 或者设置环境变量：export DEVOPS_TOKEN=<your_token>")
				fmt.Println("  3. 或者在配置文件中设置 api_token")
				os.Exit(1)
			}

			if err := auth.Login(token); err != nil {
				fmt.Fprintf(os.Stderr, "登录失败：%v\n", err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&token, "token", "t", "", "API Token")

	return cmd
}

func logoutCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "登出",
		Long:  "清除本地保存的 token",
		Run: func(cmd *cobra.Command, args []string) {
			if err := auth.Logout(); err != nil {
				fmt.Fprintf(os.Stderr, "登出失败：%v\n", err)
				os.Exit(1)
			}
		},
	}
}

func statusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "查看认证状态",
		Long:  "显示当前登录状态和用户信息",
		Run: func(cmd *cobra.Command, args []string) {
			if err := auth.Status(); err != nil {
				fmt.Fprintf(os.Stderr, "查询状态失败：%v\n", err)
				os.Exit(1)
			}
		},
	}
}
