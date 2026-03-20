package systemchange

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewSystemChangeCmd 创建 systemchange 父命令
func NewSystemChangeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "systemchange",
		Short: "系统变更管理",
		Long:  "系统变更相关操作，包括创建、列表、查看、编辑等",
	}

	cmd.AddCommand(createCmd())
	cmd.AddCommand(listCmd())
	cmd.AddCommand(viewCmd())
	cmd.AddCommand(editCmd())

	return cmd
}

func createCmd() *cobra.Command {
	var title, description, reason string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "创建系统变更",
		Long:  "创建一个新的系统变更请求",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("创建系统变更:\n")
			fmt.Printf("  标题：%s\n", title)
			fmt.Printf("  描述：%s\n", description)
			fmt.Printf("  变更原因：%s\n", reason)
			fmt.Println("\n✅ 系统变更创建成功 (示例输出)")
		},
	}

	cmd.Flags().StringVarP(&title, "title", "t", "", "变更标题 (必填)")
	cmd.Flags().StringVarP(&description, "description", "d", "", "变更描述 (必填)")
	cmd.Flags().StringVarP(&reason, "reason", "r", "", "变更原因")
	cmd.MarkFlagRequired("title")
	cmd.MarkFlagRequired("description")

	return cmd
}

func listCmd() *cobra.Command {
	var status, limit int
	var assignee string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "列出系统变更",
		Long:  "列出所有或过滤后的系统变更",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("列出系统变更:\n")
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
		Short: "查看系统变更详情",
		Long:  "查看指定系统变更的详细信息",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			fmt.Printf("查看系统变更详情:\n")
			fmt.Printf("  ID: %s\n", id)
			fmt.Println("\n✅ 查询成功 (示例输出)")
		},
	}

	return cmd
}

func editCmd() *cobra.Command {
	var title, description, status string

	cmd := &cobra.Command{
		Use:   "edit <id>",
		Short: "编辑系统变更",
		Long:  "编辑指定系统变更的信息",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			fmt.Printf("编辑系统变更:\n")
			fmt.Printf("  ID: %s\n", id)
			fmt.Printf("  新标题：%s\n", title)
			fmt.Printf("  新描述：%s\n", description)
			fmt.Printf("  新状态：%s\n", status)
			fmt.Println("\n✅ 编辑成功 (示例输出)")
		},
	}

	cmd.Flags().StringVarP(&title, "title", "t", "", "新标题")
	cmd.Flags().StringVarP(&description, "description", "d", "", "新描述")
	cmd.Flags().StringVarP(&status, "status", "s", "", "新状态")

	return cmd
}
