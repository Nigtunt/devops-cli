package main

import (
	"devops-cli/cmd"
	"devops-cli/cmd/auth"
	"devops-cli/cmd/systemchange"
	"devops-cli/cmd/task"
	"devops-cli/cmd/userstory"
	"fmt"
	"github.com/spf13/cobra"
)

// 版本信息 (通过 ldflags 注入)
var (
	Version   = "dev"
	BuildTime = "unknown"
)

func main() {
	// 注册所有子命令
	cmd.RootCmd().AddCommand(auth.NewAuthCmd())
	cmd.RootCmd().AddCommand(userstory.NewUserStoryCmd())
	cmd.RootCmd().AddCommand(systemchange.NewSystemChangeCmd())
	cmd.RootCmd().AddCommand(task.NewTaskCmd())
	
	// 添加 version 命令
	cmd.RootCmd().AddCommand(&cobra.Command{
		Use:   "version",
		Short: "显示版本信息",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("yx version %s (built: %s)\n", Version, BuildTime)
		},
	})
	
	cmd.Execute()
}
