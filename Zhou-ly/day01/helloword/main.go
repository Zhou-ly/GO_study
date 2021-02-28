package main // 声明 main 包，表明当前是一个可执行程序

import "fmt" // 导入内置 fmt 包

// 函数外只能放置标识符（变量\变量\函数\类型）的声明

func main() { // main函数，是程序执行的入口
	fmt.Println("Hello World!") // 在终端打印 Hello World!
}
