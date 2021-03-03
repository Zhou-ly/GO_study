package main

import "fmt"

func main() {

	//整型分为以下两个大类： 按长度分为：int8、int16、int32、int64 对应的无符号整型：uint8、uint16、uint32、uint64
	// uint8就是byte型，int16对应C语言中的short型，int64对应C语言中的long型。

	var i = 101

	fmt.Printf("%d\n", i)
	fmt.Printf("%b\n", i)	// 把十进制转换成二进制
	fmt.Printf("%o\n", i)	// 把十进制转换成八进制
	fmt.Printf("%x\n", i)	// 把十进制转换成十六进制

	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)	

	// 十六进制
	i3 := 0x123456
	fmt.Printf("%d\n", i3)

	// 查看变量的类型
	fmt.Printf("%T\n", i3)

	// 声明int8类型的变量
	i4 := int8(9)
	fmt.Printf("%T\n", i4)

}
