package bunch

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewHttpGatewayCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "bunch",
		Short: "bunch",
		Long:  "check and distribution",
		Run: func(cmd *cobra.Command, args []string) {
			runHttpGateway()
		},
	}
	return rootCmd
}


func runHttpGateway()  {
	fmt.Println(1111)
}