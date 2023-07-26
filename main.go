package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/debfig/bindview-cli/Tools"
	"io/ioutil"
	"net/http"
)

func main() {

	// 获取命令行输入
	state, str := Tools.GetCmd()
	if state != 1 {
		return
	}

	ui := Tools.LoadIngUi()
	// 加载动画开始
	go ui.Start()

	// 获取项目
	strURL := "https://github.com/bronze-ding/bindview-Template/archive/refs/heads/master.zip"
	res, err := http.Get(strURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	zipdata, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	zipReader, _ := zip.NewReader(bytes.NewReader(zipdata), int64(len(zipdata)))

	// 解 .zip 压缩
	err = Tools.Unzip("", zipReader)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 文件重命名
	Tools.NewFileName(&str)

	// 加载动画结束
	ui.End(str)
}
