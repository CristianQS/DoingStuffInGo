package main

import "fmt"

type book struct {
	price float64
	title string
}

func (b book) vat() float64 {
	return b.price * 0.09
}

func (b *book) discount() {
	b.price = b.price - b.price*0.1
}

func main() {
	book := book{price: 10.00, title: "The book"}
	fmt.Println(book.vat())
	book.discount()
	fmt.Println(book.price)
}
