/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"github.com/scaf-fold/gorm-gener/pkg/gener"
	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "按照配置文件生成对应表的领域模型",
	Long:  `按照配置文件采用gorm generate 完成数据表到处对应的领域模型`,
	Run:   exec,
}

func exec(cmd *cobra.Command, args []string) {
	if conf != nil {
		gener.NewModelSync(*conf).Gen()
	} else {
		panic("请查看使用，无法找到配置文件")
	}
}

var conf = new(string)

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringVarP(conf, "conf", "c", "", "-c your configuration")
}
