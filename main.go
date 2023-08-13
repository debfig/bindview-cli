package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/debfig/bindview-cli/Tools"
	"github.com/schollz/progressbar/v3"
	"io"
	"net/http"
)

// DownloadUrl 下载地址
var DownloadUrl string = "https://github.com/bronze-ding/bindview-Template/archive/refs/heads/master.zip"

// DownloadFile 文件缓存区
var DownloadFile bytes.Buffer

func main() {

	// 获取命令行输入
	state, str := Tools.GetCmd()
	if state != 1 {
		return
	}

	req, _ := http.NewRequest("GET", DownloadUrl, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"正在下载模板压缩包 (.zip)",
	)

	io.Copy(io.MultiWriter(&DownloadFile, bar), resp.Body)

	ZipData, err := io.ReadAll(&DownloadFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	ZipReader, _ := zip.NewReader(bytes.NewReader(ZipData), int64(len(ZipData)))

	// 解 .zip 压缩
	if err = Tools.Unzip("", ZipReader); err != nil {
		fmt.Println(err)
		return
	}

	// 文件重命名
	Tools.NewFileName(&str)

	// 输出命令提示
	Tools.CommandList(str)
}
