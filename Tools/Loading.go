package Tools

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Loading 加载动画结构体
type Loading struct {
	state int
	ui    []string
}

// 动画开始
func (ing Loading) Start() {
	if ing.state == 1 {
		for i := 0; i < len(ing.ui); i++ {
			time.Sleep(100 * time.Millisecond)
			cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
			cmd.Stdout = os.Stdout
			cmd.Run()
			fmt.Println("模板下载中: ", ing.ui[i])
		}
		if ing.state == 1 {
			ing.Start()
		}
	} else {
		fmt.Printf("\n")
		return
	}
}

// 动画结束
func (ing Loading) End(filename string) {
	ing.state = 0
	fmt.Println("输入下列命令启动项目:")
	fmt.Printf("\n")
	fmt.Printf("\t cd %s\n", filename)
	fmt.Printf("\t %s\n", "npm install")
	fmt.Printf("\t %s\n", "npm run serve")
	fmt.Printf("\n")
}

// 创建一个动画
func LoadIngUi() Loading {
	ui := Loading{1, []string{"|", "/", "--", "\\", "|", "/", "--", "\\"}}
	return ui
}
