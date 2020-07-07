package main

import "fmt"

func main() {

	var array [5]int

	array[0] = 100
	array[1] = 101
	array[2] = 102

	fmt.Println(array)

	for i := 0; i < len(array); i++ {
		fmt.Println(array[1])
	}

	for i, v := range array {
		fmt.Printf("Index - %v, Value - %v\n", i, v)
	}
	var arrayTwo = [...]string{"a", "b"} // массив длина которого зависит от количества элеметов в нем
	var arrayThree = [...]string{"c", "b"}

	fmt.Println(len(arrayTwo))
	fmt.Println(arrayThree == arrayTwo)

}
