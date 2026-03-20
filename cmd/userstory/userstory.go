package userstory

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewUserStoryCmd 创建 userstory 父命令
func NewUserStoryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "userstory",
		Short: "用户故事管理",
		Long:  "用户故事相关操作，包括创建、查询、更新等",
	}

	cmd.AddCommand(createCmd())
	cmd.AddCommand(listCmd())
	cmd.AddCommand(viewCmd())

	return cmd
}

func createCmd() *cobra.Command {
	var title, description string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "创建用户故事",
		Long:  "创建一个新的用户故事",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("创建用户故事:\n")
			fmt.Printf("  标题：%s\n", title)
			fmt.Printf("  描述：%s\n", description)
			fmt.Println("\n✅ 用户故事创建成功 (示例输出)")
		},
	}

	cmd.Flags().StringVarP(&title, "title", "t", "", "用户故事标题 (必填)")
	cmd.Flags().StringVarP(&description, "description", "d", "", "用户故事描述 (必填)")
	cmd.MarkFlagRequired("title")
	cmd.MarkFlagRequired("description")

	return cmd
}

func listCmd() *cobra.Command {
	var status, limit int

	cmd := &cobra.Command{
		Use:   "list",
		Short: "列出用户故事",
		Long:  "列出所有或过滤后的用户故事",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("列出用户故事:\n")
			fmt.Printf("  状态过滤：%d\n", status)
			fmt.Printf("  数量限制：%d\n", limit)
			fmt.Println("\n✅ 查询成功 (示例输出)")
		},
	}

	cmd.Flags().IntVarP(&status, "status", "s", 0, "按状态过滤")
	cmd.Flags().IntVarP(&limit, "limit", "l", 20, "返回数量限制")

	return cmd
}

func viewCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "view <id>",
		Short: "查看用户故事详情",
		Long:  "查看指定用户故事的详细信息",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id = args[0]
			fmt.Printf("查看用户故事详情:\n")
			fmt.Printf("  ID: %s\n", id)
			fmt.Println("\n✅ 查询成功 (示例输出)")
		},
	}

	return cmd
}
