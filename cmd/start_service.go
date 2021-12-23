package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serviceName string


var startServiceCmd = &cobra.Command{
	Use:   "start_service",
	Short: "start a service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start service called")
	},
}

func init() {
	startServiceCmd.Flags().StringVar(&serviceName,"name","","服务名称 (required)")
	err := startServiceCmd.MarkFlagRequired("name")
	if err != nil {
		logrus.Fatal("service name required:%v",err)
	}
	rootCmd.AddCommand(startServiceCmd)
}
