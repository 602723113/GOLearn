package main

import "fmt"

func main() {
	// 布尔类型
	var boolean bool = true
	// 数字类型
		// 整型
			var integer int = 1 // 整型
			/* 无符号 */
			var a uint8 = 255 // 无符号8位整型(0-255)
			var b uint16 = 65535 // 无符号16位整型(0-65535)
			var c uint32 = 4294967295 // 无符号16位整型(0-4294967295)
			var d uint64 = 18446744073709551615 // 无符号64位(0-18446744073709551615)
			/* 有符号 */
			var e int8 = 127 // 有符号(-128 - 127)   											-> byte
			var f int16 = 32767 // (-32768 - 32767)  											-> short
			var g int32 = 2147483647 // (-2147483648 - 2147483647)								-> int
			var h int64 = 9223372036854775807 // (-9223372036854775808 - 9223372036854775807)	-> long
		// 浮点型
		var floatOf32 float32 = 3.1415926535 // 32位浮点型数
		var floatOf64 float64 = 3.1415926535 // 64位浮点型数

	// 字符串类型
	var str string = "string"
	
	fmt.Println(boolean)
	fmt.Println(integer)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(floatOf32)
	fmt.Println(floatOf64)
	fmt.Println(str)
}