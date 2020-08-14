package main

func main() {
	// 1.
	//var a float32 = 1.1
	//var b float32 = -32.9
	//var c float64 = 44.2
	//fmt.Println(a, b, c)
	// 2.尾数部分可能丢失,造成精度损失
	//var f32 float32 = -123.000090123456789012345
	//var f64 float64 = -123.000090123456789012345
	//fmt.Println(f32, f64)
	// 3.Golang的浮点类型默认声明为float64
	//f := 3.2
	//fmt.Printf("%T\n", f)
	// 4. 浮点型可以用科学计数法表示
	//var f1 float32 = 1.23e-2
	//fmt.Println(f1)
	// 5. Golang没有专门的字符类型,作者推荐使用byte来表示字符
	// byte相当于uint8(0~255),直接输出是输出对应的码值,
	// 想要输出字符本身, 需要使用格式化输出%c
	//var c byte = 'a'
	//fmt.Println(c)
	//fmt.Printf("c = %c\n", c)
	//fmt.Printf("%T最长字节数: %d\n", c, unsafe.Sizeof(c))
	// 6.中文字符超出255范围了,不能用byte,用其他类型
	//var ch uint16 = '张'
	//fmt.Printf("ch = %c, 对应码值为%d\n", ch, ch)
	// 7. 响铃
	//fmt.Println("\a")
	// 8. 制表符\t \v
	//fmt.Println("a\tb\tc\t\n1\t2\t3\t")

}
