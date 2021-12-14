package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// startRpcCmd represents the startRpc command
var startRpcCmd = &cobra.Command{
	Use:   "startRpc",
	Short: "start rpc",
	Long:  `start rpc`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("startRpc called")
	},
}

func init() {
	rootCmd.AddCommand(startRpcCmd)
}
