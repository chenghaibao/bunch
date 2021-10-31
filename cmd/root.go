package cmd

import (
	"fmt"
	"hb_bunch/cmd/bunch"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd *cobra.Command

const Version = "1.0.0"

var Build = "local_build"

//主函数 用来调用各种网关
func init() {
	RootCmd = &cobra.Command{
		Use:              "bunch",
		Short:            "bunch",
		SilenceUsage:     true,
		SilenceErrors:    true,
		TraverseChildren: true,
		Args:             noArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return ShowHelp(os.Stderr)(cmd, args)
		},
		Version:               fmt.Sprintf("%s, build %s", Version, Build),
		DisableFlagsInUseLine: true,
	}
	RootCmd.AddCommand(bunch.NewHttpGatewayCommand())
}

//ShowHelp 查看命令行帮助.
func ShowHelp(err io.Writer) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.SetErr(err)
		cmd.HelpFunc()(cmd, args)
		return nil
	}
}

//命令行没有网关提示错误
func noArgs(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return nil
	}
	return fmt.Errorf("bunch: '%s' is not a bunch command.\nSee 'bunch --help'", args[0])
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
