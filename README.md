### ParaSnack命令行小工具

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
