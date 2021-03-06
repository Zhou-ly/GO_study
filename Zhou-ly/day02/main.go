package main

import "fmt"

// go fmt main.go	格式化代码

// Go 语言中推荐使用驼峰式命名

// var student_name string
// var studentName string
// var StudentName string

// 声明变量
// 格式
//var 变量名 变量类型

// var s1 string
// var age int
// var isOk bool

// 批量声明
var (
	name string // ''
	age  int    // 0
	isOk bool   // false
)

// 变量的初始化
// var 变量名 类型 = 表达式

func main() {
	name = "理想"
	age = 16
	isOk = true

	// Go 语言中非全局变量声明必须使用，不使用编译不过去
	// 同一个作用域({})不能重复声明同名的变量
	// 匿名变量是一个特殊的变量： _

	fmt.Print(isOk)             // 在终端中输出要打印的内容
	fmt.Println()               // 打印空行
	fmt.Printf("name:%s", name) // %s:占位符，使用name这个变量的值去替换占位符
	fmt.Println(age)            // 打印完指定的内容之后会在后面加上一个换行符

	// 声明变量同时赋值
	var s1 string = "whb"
	fmt.Println(s1)

	// 类型推导（根据值判断变量什么类型）
	var names = "ztk"
	fmt.Println(names)

	// 简短变量声明	只能在函数里面用
	s3 := "哈哈哈"
	fmt.Println(s3)

}
