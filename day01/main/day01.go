package main

import (
	"fmt"
)

func main() {
	// 1.
	// sum, sub := dualReturn(5, 1)
	// fmt.Println("sum", sum)
	// fmt.Println("sub", sub)
	// 2.
	// printForm()
	// 3.
	// variable()
	// 4.
	// var b rune = 'A'
	// fmt.Println(b)
	// 5.整型的使用细节
	// var t1 = 100 // n1是什么类型？
	// fmt.Printf("%T", t)可以做格式化输出类型
	// %v输出值
	// fmt.Printf("n1的类型 %T, n1的值%v\n", t1, t1)
	// 如何查看某个变量占用字节大小和数据类型(使用较多)
	//var n2 int64 = 10
	// unsafe.Sizeof(n) 是unsafe包里的一个函数,可以查看变量的占用字节数
	//fmt.Printf("n2的类型 %T, n2占用的字节数是%d\n", n2, unsafe.Sizeof(n2))
	//保小不保大原则(在保证程序正确运行的情况下,尽量使用占用空间小的数据类型)
	var age byte = 20
	fmt.Printf("年龄:%v, 类型:%T\n", age, age)
}

func dualReturn(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func printForm() {
	fmt.Println("name\tsex\tage\nJohn\tmale\t28")
}

var (
	n4 = "12"
	n5 = 32
	n6 = "mary"
)

func variable() {
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	var name = "tom"
	fmt.Println(name)

	//n4, n5, n6 := "sdfji", 1, 3.0
	fmt.Println(n4, n5, n6)

}
