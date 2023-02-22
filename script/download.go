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
	// 获取模板名称
	name := cCtx.Args().First()
	path := cCtx.Args().Get(1)
	url := cCtx.Args().Get(2)
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
	if url == "" {
		url = "http://172.16.0.73:12000"
	}
	// 下载模板
	var client = &http.Client{}
	url = url + "/" + name + ".tar.gz"
	fmt.Println("将要从" + url + "下载组件" + name + "到" + path)
	resp, err := client.Get(url)
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
	return nil
}
