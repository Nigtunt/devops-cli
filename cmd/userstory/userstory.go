package userstory

import (
	"fmt"

	"devops-cli/internal/pkg/format"
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
	var outputFormat string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "创建用户故事",
		Long:  "创建一个新的用户故事",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: 调用 API 创建
			result := map[string]interface{}{
				"id":     "US-123",
				"title":  title,
				"status": "created",
			}

			fmt.Printf("✅ 用户故事创建成功\n\n")
			
			// 格式化输出
			if err := formatOutput(result, outputFormat); err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "输出格式化失败：%v\n", err)
			}
		},
	}

	cmd.Flags().StringVarP(&title, "title", "t", "", "用户故事标题 (必填)")
	cmd.Flags().StringVarP(&description, "description", "d", "", "用户故事描述 (必填)")
	cmd.Flags().StringVarP(&outputFormat, "format", "f", "table", "输出格式 (table/json/yaml)")
	cmd.MarkFlagRequired("title")
	cmd.MarkFlagRequired("description")

	return cmd
}

func listCmd() *cobra.Command {
	var status, limit int
	var assignee string
	var outputFormat string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "列出用户故事",
		Long:  "列出所有或过滤后的用户故事",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: 调用 API 查询
			result := []map[string]interface{}{
				{"id": "US-123", "title": "用户登录", "status": "open"},
				{"id": "US-124", "title": "用户注册", "status": "done"},
			}

			fmt.Printf("✅ 查询成功，共 %d 条\n\n", len(result))
			
			// 格式化输出
			if err := formatOutput(result, outputFormat); err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "输出格式化失败：%v\n", err)
			}
		},
	}

	cmd.Flags().IntVarP(&status, "status", "s", 0, "按状态过滤")
	cmd.Flags().StringVarP(&assignee, "assignee", "a", "", "按负责人过滤")
	cmd.Flags().IntVarP(&limit, "limit", "l", 20, "返回数量限制")
	cmd.Flags().StringVarP(&outputFormat, "format", "f", "table", "输出格式 (table/json/yaml)")

	return cmd
}

func viewCmd() *cobra.Command {
	var outputFormat string

	cmd := &cobra.Command{
		Use:   "view <id>",
		Short: "查看用户故事详情",
		Long:  "查看指定用户故事的详细信息",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			
			// TODO: 调用 API 查询详情
			result := map[string]interface{}{
				"id":          id,
				"title":       "用户登录功能",
				"description": "作为用户，我希望能够登录系统",
				"status":      "in_progress",
				"assignee":    "张三",
			}

			fmt.Printf("✅ 查询成功\n\n")
			
			// 格式化输出
			if err := formatOutput(result, outputFormat); err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "输出格式化失败：%v\n", err)
			}
		},
	}

	cmd.Flags().StringVarP(&outputFormat, "format", "f", "table", "输出格式 (table/json/yaml)")

	return cmd
}

// formatOutput 格式化输出
func formatOutput(data interface{}, formatStr string) error {
	outputFormat, err := format.ParseFormat(formatStr)
	if err != nil {
		return err
	}
	return format.Output(data, outputFormat)
}
