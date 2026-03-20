package debug

import (
	"fmt"

	"devops-cli/internal/api"
	"github.com/spf13/cobra"
)

// NewDebugCmd 创建 debug 测试命令
func NewDebugCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "debug",
		Short: "Debug 测试命令",
		Long:  "用于测试 debug 模式的输出",
		Run: func(cmd *cobra.Command, args []string) {
			client := api.NewClient()
			
			// 如果命令行启用了 debug，确保客户端也知道
			if debug, _ := cmd.Flags().GetBool("debug"); debug {
				client.SetDebug(true)
			}
			
			fmt.Println("正在调用 /api/v1/auth/validate ...")
			
			_, err := client.ValidateToken()
			if err != nil {
				fmt.Printf("❌ 错误：%v\n", err)
				return
			}
			
			fmt.Println("✅ 请求成功")
		},
	}
}
