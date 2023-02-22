package script

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"net/http"
	"os"
)

// Download 实现下载模板到指定工作目录
func Download(cCtx *cli.Context) error {
	sync := cCtx.Bool("sync")
	name := cCtx.Args().First()
	path := cCtx.String("path")
	remote := cCtx.String("remote")
	fmt.Println("sync:" + fmt.Sprint(sync))
	fmt.Println("模板名称为:" + name + ",路径为:" + path)
	if name == "" {
		fmt.Println("请输入模板名称")
		return nil
	}
	if path == "" {
		path = "./"
	} else {
		// path文件夹是否存在，不存在则创建
		_, err := os.Stat(path)
		if err != nil {
			_ = os.MkdirAll(path, 0777)
		}
	}
	// 获取模板下载地址
	if remote == "" {
		remote = "http://172.16.0.73:12000/components"
	}
	// 下载模板
	var client = &http.Client{}
	remote = remote + "/" + name + ".tar.gz"
	fmt.Println("将要从" + remote + "下载组件" + name + "到" + path)
	resp, err := client.Get(remote)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp.StatusCode, resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	// 创建文件
	file, _ := os.Create("./" + name + ".tar.gz")
	// 写入文件
	_, _ = io.Copy(file, resp.Body)
	// 关闭文件
	_ = file.Close()
	// 解压文件到指定目录 name
	r, _ := os.Open("./" + name + ".tar.gz")
	uncompressedStream, _ := gzip.NewReader(r)
	tarReader := tar.NewReader(uncompressedStream)
	targetDir := path + "/" + name
	_ = os.MkdirAll(targetDir, 0777)
	for true {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// 获取文件信息
		target := targetDir + "/" + header.Name
		info := header.FileInfo()
		if info.IsDir() {
			_ = os.MkdirAll(target, info.Mode())
			continue
		}
		// 创建文件
		if info.Mode().IsRegular() {
			outFile, _ := os.Create(target)
			_, _ = io.Copy(outFile, tarReader)
			_ = outFile.Close()
		}
	}
	// 删除压缩包
	_ = os.Remove("./" + name + ".tar.gz")
	// 同步创建Snack组件文件并引用
	if sync {
		err := Generate(cCtx)
		if err != nil {
			return err
		}
	}
	return nil
}
