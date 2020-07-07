package main

import "fmt"

func main() {

	t := calculate(100)

	//t := func(k int) int {
	//	return k * 1000
	//}

	z := t(5)
	fmt.Println(z)
}

func calculate(i int) func(int) int {

	fmt.Println(i)

	return func(k int) int {
		return k * 1000
	}
}
