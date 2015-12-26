
Write a simple lexer which allows you to run simple SQL-like queries against
a slice of structs. For example, if you have this data:

```
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
     Status    Status}
```

From there we should be able to write queries like:

```
SELECT * FROM Candidate
SELECT FirstName, Age FROM Candidate
```

Bonus if you also spend the time to get filters working:

```
SELECT * FROM Candidate WHERE Status = Applied
```

Altogether, the library should work something like this:

```
func main() {
  db, err := InitDB()
  if err != nil {
    panic(err)
  }
  for _, c in db.Query("SELECT * FROM Candidate") {
    fmt.Printf("Retrieved candidate: %v", c)
  }
}
```

## References

1. [See Rob Pike's talk about Lexical Scanning in Go](https://www.youtube.com/watch?v=HxaD_trXwRE)
2. [Here is a nice tutorial on writing a parser/lexer in Go.](https://blog.gopheracademy.com/advent-2014/parsers-lexers/)

## Implementations

After you've taken a stab, here are some reference implementations:

1. [sql.go](./sql.go) - simple with writing a custom lexer in Go