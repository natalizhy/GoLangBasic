package main

import "fmt"

func main() {
	var sliceOne []int
	fmt.Println(len(sliceOne))
	fmt.Println(cap(sliceOne))

	sliceOne = append(sliceOne, 10)
	sliceOne = append(sliceOne, 20)
	sliceOne = append(sliceOne, 30)

	fmt.Println(len(sliceOne))
	fmt.Println(cap(sliceOne))

	fmt.Println(sliceOne[:3]) //режем слайс

	slice := make([]int, 5)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice[4])
	fmt.Println(cap(sliceOne))

	x := []int{1, 2, 3}

	fmt.Println(x)
}
