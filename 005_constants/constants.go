package main

import "fmt"

//const applicationName = "App Const"
const constMySpecial = "App Const"

const (
	_ = iota // автоматическое  инкрементарное значение начинаеться с 0
	_ = iota
	_ = iota
	D = iota + 5
)

func main() {

	var temp string = "simple string"
	fmt.Println(temp)

	//fmt.Println(A)
	//fmt.Println(B)
	//fmt.Println(C)
	fmt.Println(D)

	//const number = 100
	//fmt.Println(applicationName)
	//fmt.Println(number)
	//
	//var numberTwo int = 10
	//result := number + numberTwo
	//fmt.Println(result)

	x, y, z := TypeGo()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}

func TypeGo() (int, bool, string)  {
	return 100, false, "stringType"
}
