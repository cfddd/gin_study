package main

import (
	"fmt"
)

const (
	cBlock  = 0
	cRed    = 1
	cGreed  = 2
	cYellow = 3
	cBlue   = 4
	cPurple = 5
	cCyan   = 6
	cGrey   = 7
)

func PrintColor(colorCode int, text string, isBackGround bool) {
	if !isBackGround {
		fmt.Printf("\033[3%dm%s\033[0m", colorCode, text)
	} else {
		fmt.Printf("\033[4%dm%s\033[0m", colorCode, text)
	}
}

func main() {
	fmt.Println("\033[29m白色\033[0m")
	fmt.Println("\033[30m黑色\033[0m")
	fmt.Println("\033[31m红色\033[0m")
	fmt.Println("\033[32m绿色\033[0m")
	fmt.Println("\033[33m黄色\033[0m")
	fmt.Println("\033[34m蓝色\033[0m")
	fmt.Println("\033[35m紫色\033[0m")
	fmt.Println("\033[36m青色\033[0m")
	fmt.Println("\033[37m灰色\033[0m")
	//有颜色的只有这么多
	//for i := 0; i < 100; i++ {
	//	fmt.Println("\033[" + strconv.Itoa(i) + "m灰色\033[0m" + strconv.Itoa(i))
	//}

	//背景色
	fmt.Println("\033[40n黑色\033[0m")
	fmt.Println("\033[41m红色\033[0m")
	fmt.Println("\033[42m绿色\033[0m")
	fmt.Println("\033[43m黄色\033[0m")
	fmt.Println("\033[44m蓝色\033[0m")
	fmt.Println("\033[45m紫色\033[0m")
	fmt.Println("\033[46m青色\033[0m")
	fmt.Println("\033[47m灰色\033[0m")
}
