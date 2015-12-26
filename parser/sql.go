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

var Candidates = []Candidate{
	Candidate{"Jack", "Doe", 25, "555-", 0},
	Candidate{"Jill", "Doe", 30, "555-", 1},
	Candidate{"Jack", "Murphy", 35, "555-", 2},
	Candidate{"Jill", "Murphy", 45, "555-", 3},
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

func (l *Lexer) Items() []Item {
	slice := make([]Item, 0)
	for item := range l.items {
		slice = append(slice, item)
	}
	return slice
}

type SelectQuery struct {
	Table string
	Fields []string
	AllFields bool
}


func ItemPosition(it ItemType, items []Item) (int, bool) {
	for i, val := range items {
		if val.ItemType == it {
			return i, true
		}
	}
	return 0, false
}

func ParseSelectQuery(items []Item) (SelectQuery, error) {
	sq := SelectQuery{}

	fromPos, foundFrom := ItemPosition(FromItem, items)
	if !foundFrom {
		return sq, fmt.Errorf("Didn't find any FROM")
	}

	// get the fields to select
	fields := make([]string, 0)
	for _, item := range items[:fromPos] {
		switch item.ItemType {
		case AsterixItem:
			// use all fields
			sq.AllFields = true
		case StringItem:
			field := strings.Trim(item.Content, ", ")
			fields = append(fields, field)
		default:
			return sq, fmt.Errorf("%v was neither asterix nor a string", item.String())
		}
	}
	sq.Fields = fields

	if (fromPos + 3) > len(items) {
		return sq, fmt.Errorf("must specify a table to select from.")

	}

	tableItem := items[fromPos+1]
	if tableItem.ItemType != StringItem {
		return sq, fmt.Errorf("%v is not a string", tableItem.String())
	}

	sq.Table = tableItem.Content
	return sq, nil
}

func Query(q string) ([]Candidate, error) {
	cands := make([]Candidate, 0)
	l := NewLexer(q)

	items := l.Items()
	if len(items) < 1 {
		return cands, fmt.Errorf("invalid query (too short)")
	}

	switch items[0].ItemType {
	case SelectItem:
		sq, err := ParseSelectQuery(items[1:])
		if err != nil {
			return cands, err
		}
		if sq.Table != "Candidate" {
			return cands, fmt.Errorf("table %v was not one of [Candidates]", sq.Table)
		}
		if sq.AllFields {
			sq.Fields = []string{"FirstName", "LastName", "Age", "Phone", "Status"}
		}

		// reflect would be more elegant, albeit slower, here
		for _, cand := range Candidates {
			c := Candidate{}
			for _, field := range sq.Fields {
				switch field {
				case "FirstName":
					c.FirstName = cand.FirstName
				case "LastName":
					c.LastName = cand.LastName
				case "Age":
					c.Age = cand.Age
				case "Phone":
					c.Phone = cand.Phone
				case "Status":
					c.Status = cand.Status
				}
			}
			cands = append(cands, c)
		}

		return cands, nil

	default:
		err := fmt.Errorf("can't handle query of type: %v", items[0].Content)
		return cands, err
	}
}



func main() {
	queries := []string{"SELECT * FROM Candidate", "SELECT FirstName, LastName FROM Candidate", "UPDATE * FROM Candidate", "SELECT a, b, c FROM"}
	for _, query := range queries {
		fmt.Println(query)
		resp, err := Query(query)
		if err != nil {
			fmt.Printf("\terror querying: %v\n", err)
		}
		fmt.Printf("\tcandidates: %v\n", resp)
	}
}
