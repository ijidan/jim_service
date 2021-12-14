package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// startWsCmd represents the startWs command
var startWsCmd = &cobra.Command{
	Use:   "startWs",
	Short: "start websocket",
	Long:  `start websocket`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("startWs called")
	},
}

func init() {
	rootCmd.AddCommand(startWsCmd)
}
