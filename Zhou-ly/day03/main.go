package main

import "fmt"

// 常量
// 定义了常量之后不能改变
// 在程序运行期间不会改变的量

const pi = 3.1415926

// 批量声明常量
const (
	statusOK = 200
	notFOUND = 404
)

// 批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一致
const (
	n1 = 100
	n2
	n3
)

// iota是go语言的常量计数器，只能在常量的表达式中使用。
// iota在const关键字出现时将被重置为0。
// const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 
// 使用iota能简化定义，在定义枚举时很有用。

const (
	n4 = iota //0
	n5        //1
	n6        //2
	n7        //3
)

// 几个常见的iota示例:

// 使用_跳过某些值
// const (
// 	n1 = iota //0
// 	n2        //1
// 	_
// 	n4        //3
// )

// iota声明中间插队
// const (
// 	n1 = iota //0
// 	n2 = 100  //100
// 	n3 = iota //2
// 	n4        //3
// )
// const n5 = iota //0

// 定义数量级 这里的<<表示左移操作
// const (
// 	_  = iota
// 	KB = 1 << (10 * iota)
// 	MB = 1 << (10 * iota)
// 	GB = 1 << (10 * iota)
// 	TB = 1 << (10 * iota)
// 	PB = 1 << (10 * iota)
// )

// 多个iota定义在一行
// const (
// 	a, b = iota + 1, iota + 2 //1,2
// 	c, d                      //2,3
// 	e, f                      //3,4
// )

func main(){
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	fmt.Println("n4:", n4)
	fmt.Println("n5:", n5)
	fmt.Println("n6:", n6)
	fmt.Println("n7:", n7)
}
