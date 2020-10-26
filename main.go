package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"yuemiao/config"
	"yuemiao/jin_niu"
	"yuemiao/seckill"
	"yuemiao/yuemiao"
)

var (
	cfgFile string
)

func yuemiaoCmd() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "yuemiao",
		Short: "启动约苗九价脚本",
		Long:  "启动约苗九价脚本",
		RunE: func(cmd *cobra.Command, args []string) error {
			s := yuemiao.NewYueMiao(zap.L(), config.C.YueMiao)
			s.V2()
			zap.L().Info("yuemiao")
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			config.Init(cfgFile)
			return nil
		},
	}

	serverCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	return serverCmd
}

func vcodeCmd() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "vcode",
		Short: "验证码相关命令",
		Long:  "验证码相关命令",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			config.Init(cfgFile)
			return nil
		},
	}

	serverCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	serverCmd.AddCommand(getVcodeCmd())
	serverCmd.AddCommand(parseVcodeCmd())
	return serverCmd
}

func getVcodeCmd() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "get",
		Short: "获取所有验证码",
		Long:  "获取所有验证码",
		Run: func(cmd *cobra.Command, args []string) {
			s := yuemiao.NewYueMiao(zap.L(), config.C.YueMiao)
			err := s.GetAllVCode()
			if err != nil {
				fmt.Printf("获取所有验证码失败：%s\n", err.Error())
			}
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			config.Init(cfgFile)
			return nil
		},
	}
	return serverCmd
}

func parseVcodeCmd() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "parse",
		Short: "解析所有验证码",
		Long:  "解析所有验证码",
		Run: func(cmd *cobra.Command, args []string) {
			s := yuemiao.NewYueMiao(zap.L(), config.C.YueMiao)
			err := s.ParseAllVCode()
			if err != nil {
				fmt.Printf("解析所有验证码失败：%s\n", err.Error())
			}
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			config.Init(cfgFile)
			return nil
		},
	}
	return serverCmd
}

func yuemiaoSeckillCmd() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "seckill",
		Short: "查看约苗有哪些城市有秒杀信息",
		Long:  "查看约苗有哪些城市有秒杀信息",
		RunE: func(cmd *cobra.Command, args []string) error {
			config.C.YueMiao.Verbose = false
			s := seckill.NewAllSteps(zap.L(), config.C.YueMiao)
			cities, err := s.GetSeckillCities()
			if err != nil {
				fmt.Printf("查看约苗有哪些城市有秒杀信息有误：%s\n", err.Error())
			}
			fmt.Printf("明天可以秒杀的城市有：%s\n", cities)

			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			config.Init(cfgFile)
			return nil
		},
	}

	serverCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	return serverCmd
}

func jinniuCmd() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "jinniu",
		Short: "启动金牛公众号九价脚本",
		Long:  "启动金牛公众号九价脚本",
		RunE: func(cmd *cobra.Command, args []string) error {
			zap.L().Info("jinniu")
			j := jin_niu.NewJinNiu(zap.L(), config.C.JinNiu)
			j.Together()
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			config.Init(cfgFile)
			return nil
		},
	}

	serverCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	return serverCmd
}

var RootCmd cobra.Command = cobra.Command{
	Use: "jiujia",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func main() {
	RootCmd.AddCommand(yuemiaoCmd())
	RootCmd.AddCommand(jinniuCmd())
	RootCmd.AddCommand(yuemiaoSeckillCmd())
	// RootCmd.AddCommand(vcodeCmd())
	if err := RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
