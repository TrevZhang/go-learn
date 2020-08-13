package main

import "fmt"

func main() {
	// 1.
	// sum, sub := dualReturn(5, 1)
	// fmt.Println("sum", sum)
	// fmt.Println("sub", sub)
	// 2.
	printForm()
}

func dualReturn(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func printForm() {
	fmt.Println("name\tsex\tage\nJohn\tmale\t28")
}
