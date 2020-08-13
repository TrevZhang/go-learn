package main

import "fmt"

func main() {
	// 1.
	// sum, sub := dualReturn(5, 1)
	// fmt.Println("sum", sum)
	// fmt.Println("sub", sub)
	// 2.
	//printForm()
	// 3.
	variable()
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
