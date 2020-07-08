package main

import "fmt"

type EDocument struct {
	Number       string
	Date         string
	NumberOPages int
	Footer
}

type Footer struct {
	Author string
}

func main() {

	doc1 := EDocument{
		Number:       "001-A",
		Date:         "08.07.2020",
		NumberOPages: 10,
		Footer: Footer{
			Author: "Author-1",
		},
	}

	var doc2 EDocument
	doc2.Number = "002-A"
	doc2.Date = "08.07.2020"
	doc2.NumberOPages = 5
	doc2.Footer.Author = "Author-2"


	fmt.Printf("%T - %v\n", doc1, doc1)
	fmt.Printf("%T - %v\n", doc2, doc2)
	fmt.Printf("%T - %v\n", doc2.Footer, doc2.Footer)

}
