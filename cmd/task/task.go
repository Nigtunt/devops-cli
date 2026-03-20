package task

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewTaskCmd 创建 task 父命令
func NewTaskCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task",
		Short: "任务管理",
		Long:  "任务相关操作，包括创建、列表、查看等",
	}

	cmd.AddCommand(createCmd())
	cmd.AddCommand(listCmd())
	cmd.AddCommand(viewCmd())

	return cmd
}

func createCmd() *cobra.Command {
	var title, description, assignee string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "创建任务",
		Long:  "创建一个新的任务",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("创建任务:\n")
			fmt.Printf("  标题：%s\n", title)
			fmt.Printf("  描述：%s\n", description)
			fmt.Printf("  负责人：%s\n", assignee)
			fmt.Println("\n✅ 任务创建成功 (示例输出)")
		},
	}

	cmd.Flags().StringVarP(&title, "title", "t", "", "任务标题 (必填)")
	cmd.Flags().StringVarP(&description, "description", "d", "", "任务描述")
	cmd.Flags().StringVarP(&assignee, "assignee", "a", "", "负责人")
	cmd.MarkFlagRequired("title")

	return cmd
}

func listCmd() *cobra.Command {
	var status, limit int
	var assignee string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "列出任务",
		Long:  "列出所有或过滤后的任务",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("列出任务:\n")
			fmt.Printf("  状态过滤：%d\n", status)
			fmt.Printf("  负责人：%s\n", assignee)
			fmt.Printf("  数量限制：%d\n", limit)
			fmt.Println("\n✅ 查询成功 (示例输出)")
		},
	}

	cmd.Flags().IntVarP(&status, "status", "s", 0, "按状态过滤")
	cmd.Flags().StringVarP(&assignee, "assignee", "a", "", "按负责人过滤")
	cmd.Flags().IntVarP(&limit, "limit", "l", 20, "返回数量限制")

	return cmd
}

func viewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view <id>",
		Short: "查看任务详情",
		Long:  "查看指定任务的详细信息",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			fmt.Printf("查看任务详情:\n")
			fmt.Printf("  ID: %s\n", id)
			fmt.Println("\n✅ 查询成功 (示例输出)")
		},
	}

	return cmd
}
