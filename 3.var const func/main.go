package main

import "fmt"

func main() {
	// var 可以修饰一个量为常量
	var a = "1"
	fmt.Println(a)
	a = "2"
	fmt.Println(a)

	fmt.Println("======")

	// const 可以修饰该量为常量
	const ONE = "1"
	fmt.Println(ONE)

	fmt.Println("======")


	// func
	sendMessage("测试")

	x,y := swapValue(1, 2)
	fmt.Println(x, y)
}

/*
	func 可以用于声明函数
	func 函数名(参数) 返回值 {
   		函数体
	}	
 */
func sendMessage(message string) {
	fmt.Println(message)
}


/*
	在go-lang中函数可以有多返回值
	例如
 */
func swapValue(a b int) (int int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}


/*
	直接返回两个不同的顺序
 */
func swapValue2(a b int) (int int) {
	return b, a
}