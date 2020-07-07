package main

import "fmt"

func main()  {

	var simpleMap = map[int]string{}
	fmt.Println(simpleMap)
	simpleMap[198] = "some word"

	fmt.Println(simpleMap)

	var anotherMap = make(map[string]int)
	anotherMap["one"] = 198
	anotherMap["two"] = 1981
	anotherMap["three"] = 1982

	fmt.Println(anotherMap)

	delete(anotherMap, "three") // удаление по ключу
	fmt.Println(anotherMap["two"])

	slice := []int{1, 2, 4}

	for _, v := range slice {
		fmt.Println(v)
		if v ==2 {
			fmt.Println("ok")
		}
	}

	if value, ok := anotherMap["one"]; ok {  // по ключу есть значение
		fmt.Println(value)
	}
}
