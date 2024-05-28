## ParaSnack命令行小工具

基于开源协议，如果觉得不错，可以动手给个小心心❤️,[仓库直达](https://github.com/linktomarkdown/para-cli)

ParaSnack是一个命令行小工具，用于快速生成一个模板

```bash
para -h
```

#### 安装

```bash
curl -o- https://raw.githubusercontent.com/linktomarkdown/para-cli/main/install.sh | bash
```


#### 使用

##### 上传组件
```bash
para upload/up "文件夹路径"
```

##### 下载组件
```bash
para get/g "组件名" # 拉取组件到当前目录，路径为./

para get/g -s/sync  "组件名" # -s 是否同步创建Snack组件

para get/g -p "./src/components" Foo # 拉取组件到指定目录，路径为./src/components/Foo (前提，目录存在)

para get/g -p "./src/components" -s/sync Foo # 拉取组件到指定目录，路径为./src/components/Foo (前提，目录存在)，并且同步创建Snack组件
```

##### 创建Snack模板
```bash
para new/n -p "./src" Demo # 创建Snack模板到指定目录，路径为./src/Demo (前提，目录(src)存在)
```

# Para-cli 手动安装脚本安装步骤

## 程序安装

1. 下载程序 [下载传送门](https://github.com/linktomarkdown/para-cli/releases)

## MacOS

2. 程序安装

- 将下载包中的 para-cli 程序放到`/usr/local/bin`目录下

- 软链接到`para`, 可以用`para`作为命令访问，在终端中以管理员执行如下命令回车输入电脑密码即可

```bash
sudo ln -s /usr/local/bin/para-cli /usr/local/bin/para
```

3. 结果验证

- `para -h`

```
➜  bin para -h
NAME:
   ParaCLI - ParaSnack模板生成脚手架!

USAGE:
   ParaCLI [global options] command [command options] [arguments...]

VERSION:
   0.1.8

COMMANDS:
   get, g      拉取仓库组件到本地Snack工程下的Components目录. -s 是否同步创建Snack组件文件并引用. -r 指定生成Snack组件文件的地址. -p 指定拉取组件的路径 -pp 指定拉取组件的页面.
   new, n      生成新的Snack模板. p 指定拉取组件的路径.
   upload, up  提交组件到仓库. -p 文件夹路径.
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --get value, -g value       拉取仓库组件到本地Snack工程下的Components目录.
   --new value, -n value       生成新的Snack模板.
   --upload value, --up value  上传组件到仓库.
   --help, -h                  show help
   --version, -v               print the version
```

## windows 配置环境变量

> 以 windows 10 为例

- 桌面右键**我的电脑(也可能叫此电脑)**，点击属性，也可以从系统设置进入，找到**关于**
- 选择高级系统设置
- 选择环境变量
- 选择系统变量里的 Path，点击编辑
- 点击新建，填入文件路径即可，注意不用带文件名

