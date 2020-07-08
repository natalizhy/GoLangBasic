package main

import "fmt"

type Document struct { // структуры данных
	Date          string
	Number        string
	NumberOfPages int
}

type PersonCard struct {
	Date      string
	FirstName string
	LastName  string
	Age       int
}

type PrintInterface interface {
	CheckData()
}

func (d *Document) CheckData() { // метод
	if d.Date != "" {
		fmt.Println("Date in the doc - correct")
	} else {
		fmt.Println("Date in the doc - not correct")
	}
}
func (d *PersonCard) CheckData() {
	if d.Date != "" {
		fmt.Println("Date in the person card - correct")
	} else {
		fmt.Println("Date in the person card - not correct")
	}
}

func main() {

	doc := new(Document)
	pcard := new(PersonCard)

	doc.Date = "8.07.2020"
	doc.NumberOfPages = 5
	doc.Number = "A - 100"

	pcard.Date = "8.07.2020"
	pcard.Age = 21
	pcard.FirstName = "User"
	pcard.LastName = "Test"

	sl := []PrintInterface{doc, pcard}

	PrintAnyDoc(sl)
}

func PrintAnyDoc(anyDoc []PrintInterface) {
	for _, v := range anyDoc {
		fmt.Println(v)
	}
}
