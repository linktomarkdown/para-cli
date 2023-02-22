package script

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func Generate(cCtx *cli.Context) error {
	// 获取模板名称
	name := cCtx.Args().First()
	// 生成模板文件的路径，默认为当前目录
	path := cCtx.Args().Get(1)
	if name == "" {
		fmt.Println("请输入模板名称")
		return nil
	}
	if path == "" {
		path = "./"
	}
	fmt.Printf("生成模板名称为:%s,生成路径为:%s\n", name, path)
	// 生成模板的命令行工具,将SelfcareTemplate目录下的模板文件复制到指定目录,并将模板文件中的模板名称替换为指定的名称
	// 创建文件夹 name/src
	rootDir := path + "/" + name
	if os.Mkdir(rootDir, 0777) != nil {
		fmt.Println("创建文件夹失败")
		return nil
	}
	srcPath := rootDir + "/src"
	if os.Mkdir(srcPath, 0777) != nil {
		fmt.Println("创建文件夹失败")
		return nil
	}
	stylePath := srcPath + "/style"
	if os.Mkdir(stylePath, 0777) != nil {
		fmt.Println("创建文件夹失败")
		return nil
	}
	// 创建文件 name/src/style/index.scss
	if _, err := os.Create(stylePath + "/index.scss"); err != nil {
		fmt.Println("创建文件失败")
		return nil
	}
	// name/src/index.tsx,写入默认数据
	defIndexData := `import './style/index.scss';
import React from 'react';
import { Snack } from '@para-snack/core';
import ParauiProvider from '@para-ui/core/ParauiProvider';
import local from '../../../locale';
interface Props {
}
export class ` + name + ` extends Snack {
	constructor(data: Props) {
		super(data);
		this.$i18n = local;
	}

	public $component(): JSX.Element {
		return (
			<ParauiProvider seed={'` + name + `'} productionPrefix={'` + name + `'}>
				
			</ParauiProvider>
		);
	}
}
export default ` + name + `;`
	if err := os.WriteFile(srcPath+"/index.tsx", []byte(defIndexData), 0777); err != nil {
		fmt.Println("创建文件失败")
		return nil
	}
	// name/src/setting.tsx,写入默认数据
	defSettingData := `import React from 'react';
import {
  Snack,
  SnackSetting,
} from '@para-snack/core';
interface Props {}
export class ` + name + "Setting" + ` extends SnackSetting {
	constructor(public main: Snack, public data: Props) {
		super(data);
	}
	public $component() {
		return (<></>);
	}
}`
	if err := os.WriteFile(srcPath+"/setting.tsx", []byte(defSettingData), 0777); err != nil {
		fmt.Println("创建文件失败")
		return nil
	}
	// 创建文件 name/index.ts，写入默认数据
	defRootData := `export * from "./src/index";
export * from "./src/setting";`
	if err := os.WriteFile(rootDir+"/index.ts", []byte(defRootData), 0777); err != nil {
		fmt.Println("创建文件失败")
		return nil
	}
	defJsonData := `{"name": "` + name + `", "description": "` + name + `", "version": "1.0.0.1", "author": "ParaCLI"}`
	// 创建name/snack.json 文件，写入默认数据
	if err := os.WriteFile(rootDir+"/snack.json", []byte(defJsonData), 0777); err != nil {
		fmt.Println("创建文件失败")
		return nil
	}
	return nil
}
