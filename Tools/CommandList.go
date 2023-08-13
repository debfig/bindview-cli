package Tools

import (
	"fmt"
	"github.com/zhangboqi/cfmt"
)

// CommandList 命令列表
func CommandList(filename string) {
	// 换行
	fmt.Printf("\n")
	// green
	cfmt.Yprintf("\t 输入下列命令启动项目: \n")
	// 换行
	fmt.Printf("\n")
	// red
	cfmt.Rprintf("\t cd %s\n", filename)
	cfmt.Rprintf("\t npm install\n")
	cfmt.Rprintf("\t npm run serve\n")
	// 换行
	fmt.Printf("\n")
}
