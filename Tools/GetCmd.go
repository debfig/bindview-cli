package Tools

import (
	"fmt"
	"os"
)

// GetCmd 获取命令行输入
func GetCmd() (int, string) {

	var parame []string

	for _, val := range os.Args[1:] {
		parame = append(parame, val)
	}

	if len(parame) == 0 {
		fmt.Println("使用 -h 查看命令")
		return 0, ""
	}

	switch parame[0] {
	case "-h":
		fmt.Printf("用法\n")
		fmt.Printf("\t\n < -v > 显示版本")
		fmt.Printf("\t\n < create > 创建一个项目, 例: create myapp")
		return 0, ""
	case "-v":
		fmt.Println("v1.1.0")
		return 0, ""
	case "create":
		if len(parame) == 2 {
			return 1, parame[1]
		} else {
			fmt.Println("create 命令需要一个文件名")
			return 0, ""
		}
	default:
		fmt.Println("使用 -h 查看命令")
		return 0, ""
	}
}
