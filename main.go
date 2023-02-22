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
		Name:                   "ParaCLI",
		Usage:                  "ParaSnack模板生成脚手架!",
		Version:                "0.1.7",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "生成新的Snack模板.",
			},
			&cli.StringFlag{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "拉取仓库组件到本地Snack工程下的Components目录.",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "生成新的Snack模板. -s 是否同步创建Snack组件文件并引用. -r 指定生成Snack组件文件的地址. -p 指定拉取组件的路径 -pp 指定拉取组件的页面.",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "sync", Aliases: []string{"s"}},
					&cli.StringFlag{Name: "path", Aliases: []string{"p"}},
					&cli.StringFlag{Name: "remote", Aliases: []string{"r"}},
					&cli.StringFlag{Name: "ppath", Aliases: []string{"pp"}},
				},
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
				Usage:   "拉取仓库组件到本地Snack工程下的Components目录. -s 是否同步创建Snack组件文件并引用. -r 指定生成Snack组件文件的地址. -p 指定拉取组件的路径 -pp 指定拉取组件的页面.",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "sync", Aliases: []string{"s"}},
					&cli.StringFlag{Name: "path", Aliases: []string{"p"}},
					&cli.StringFlag{Name: "remote", Aliases: []string{"r"}},
					&cli.StringFlag{Name: "ppath", Aliases: []string{"pp"}},
				},
				Action: func(cCtx *cli.Context) error {
					//fmt.Println("sync:", cCtx.Bool("sync"))
					//fmt.Println("path:", cCtx.String("path"))
					//fmt.Println("remote:", cCtx.String("remote"))
					//fmt.Println("name:", cCtx.Args().First())
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
