package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"para/script"
	"sort"
)

func main() {
	app := &cli.App{
		Name:    "ParaCLI",
		Usage:   "ParaSnack模板生成脚手架!",
		Version: "0.1.6",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "快速生成新的Snack模板",
			},
			&cli.StringFlag{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "快速获取Snack 组件",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "生成新的Snack模板",
				Action: func(cCtx *cli.Context) error {
					err := script.Generate(cCtx)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "拉取仓库组件到本地Snack工程下的Components目录",
				Action: func(cCtx *cli.Context) error {
					err := script.Download(cCtx)
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
