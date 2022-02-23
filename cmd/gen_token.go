package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"jim_service/pkg"
)
var userId uint64


// genGormCmd represents the genGorm command
var genTokenCmd = &cobra.Command{
	Use:   "gen_token",
	Short: "gen token",
	Run: func(cmd *cobra.Command, args []string) {
		token:=pkg.GenJwtToken(userId,pkg.Conf.Jwt.Secret)
		fmt.Printf( "token generated：%s",token)
	},
}

func init() {
	genTokenCmd.Flags().Uint64Var(&userId,"id",0,"用户ID (required)")
	err := genTokenCmd.MarkFlagRequired("id")
	if err != nil {
		logrus.Fatal("flat required:%v",err)
	}
	rootCmd.AddCommand(genTokenCmd)

}
