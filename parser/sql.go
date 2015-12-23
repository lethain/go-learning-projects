package main

import (
	"fmt"
	"strings"
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



type ItemType string
const (
	StringItem ItemType = "String"
	SelectItem ItemType = "Select"
	FromItem ItemType = "From"
	WhereItem ItemType = "Where"
	AsterixItem ItemType = "Asterix"
	EOFItem ItemType = "EOF"
)

type Item struct {
	ItemType ItemType
	Content string
}

func (i *Item) String() string {
	switch i.Content {
	case "":
		return fmt.Sprintf("Item(%v)", i.ItemType)
	default:
		return fmt.Sprintf("Item(%v, \"%v\")", i.ItemType, i.Content)
	}
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
	remaining := l.input
	for _, word := range strings.Split(remaining, " ") {
		switch strings.ToLower(word) {
		case "select":
			l.items <- Item{ItemType: SelectItem}
		case "from":
			l.items <- Item{ItemType: FromItem}
		case "where":
			l.items <- Item{ItemType: WhereItem}
		case "*":
			l.items <- Item{ItemType: AsterixItem}				
		default:
			l.items <- Item{ItemType: StringItem, Content: word}
		}
	}
	l.items <- Item{ItemType: EOFItem}
	close(l.items)
}


func (l *Lexer) Items() chan Item {
	return l.items
}

func main() {
	i := "SELECT * FROM Candidate"
	l := NewLexer(i)
	fmt.Println(l)
	for item := range l.Items() {
		fmt.Println(item.String())
	}
}
