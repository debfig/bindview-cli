package Tools

import (
	"fmt"
	"os"
)

// NewFileName 重命名文件夹
func NewFileName(newfilename *string) {
	from := "./bindview-Template-master"
	var to string

	// 判断目录是否存在
	_, err := os.Stat(*newfilename)
	if err == nil {
		fmt.Println("目录已存在,将使用默认名字")
		*newfilename = "bindview-Template-master"
		return
	}

	if *newfilename == "" {
		to = "bindview-template"
	} else {
		to = *newfilename
	}

	err = os.Rename(from, to)
	if err != nil {
		fmt.Println(err)
		return
	}
}
