package script

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func Upload(cCtx *cli.Context) error {
	path := cCtx.Args().First()
	//remote := "http://localhost:13000/uploadFile"
	remote := "http://172.16.0.73:13000/uploadFile"
	if path == "" {
		fmt.Println("请输入上传的文件路径")
		return nil
	}
	// 名称为路径最后一个文件夹的名字
	name := filepath.Base(path)
	fmt.Println("模板名称为:" + name + ",路径为:" + path)
	// 判断是否是文件夹
	_, err := os.Stat(path)
	if err != nil {
		fmt.Println("路径不存在")
		return nil
	}
	// 压缩文件夹为tar.gz
	var buf bytes.Buffer
	_ = Compress(path, &buf)
	// write the .tar.gzip
	fileToWrite, _ := os.OpenFile("./"+name+".tar.gz", os.O_CREATE|os.O_RDWR, os.FileMode(777))
	if _, err := io.Copy(fileToWrite, &buf); err != nil {
		log.Fatal(err)
		return nil
	}
	// 上传文件 remote
	_ = UploadFile("./"+name+".tar.gz", remote)
	// 上传成功后删除文件
	_ = os.Remove("./" + name + ".tar.gz")
	return nil
}

func Compress(src string, buf io.Writer) error {
	// tar > gzip > buf
	zr := gzip.NewWriter(buf)
	tw := tar.NewWriter(zr)
	// walk through every file in the folder
	_ = filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {
		// generate tar header
		header, _ := tar.FileInfoHeader(fi, file)
		// must provide real name
		// (see https://golang.org/src/archive/tar/common.go?#L626)
		header.Name = filepath.ToSlash(file)
		// write header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		// if not a dir, write file content
		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			if _, err := io.Copy(tw, data); err != nil {
				return err
			}
		}
		return nil
	})
	// produce tar
	if err := tw.Close(); err != nil {
		return err
	}
	// produce gzip
	if err := zr.Close(); err != nil {
		return err
	}
	//
	return nil
}

func UploadFile(path string, url string) error {
	//curl -X POST http://localhost:8080/upload \
	//-F "file=@/Users/appleboy/test.zip" \
	//-H "Content-Type: multipart/form-data"
	// 上传文件
	file, _ := os.Open(path)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func(resp *http.Response) {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp)
	fmt.Println(resp.StatusCode)
	return nil
}
