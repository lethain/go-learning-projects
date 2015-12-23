package main

import (
	"fmt"
)

type Status int

const (
      Applied     Status = iota
      PhoneScreen
      Onsite
      Offer
      Accepted
      Rejected
)

type Candidate struct {
     FirstName string
     LastName  string
     Age       int
     Phone     string
     Status    Status
}



type ItemType int
const (
	SelectItem ItemType = iota
	ForItem
	WhereItem
	AsterixItem
	EOFItem
)

type Item struct {
	ItemType ItemType
	Content string
}

type Lexer struct {
	input string
	items chan Item
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.items = make(chan Item)
	go l.Parse()
	return l
}

func (l *Lexer) String() string {
	return fmt.Sprintf("Lexer(%v)", l.input)
}

func (l *Lexer) Parse() {
	fmt.Printf("%v - parsing finished\n", l)
	l.items <- Item{ItemType: EOFItem}
	close(l.items)
}


func (l *Lexer) Items() chan Item {
	return l.items

}

func main() {
	i := "SELECT * FROM Candidate"	
	l := NewLexer(i)
	fmt.Printf("NewLexer(%v)\n", l)
	for item := range l.Items() {
		fmt.Printf("Item: %v\n", item)
	}
}
