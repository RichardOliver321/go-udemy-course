# go-udemy-course

## Package

- workspace
- related code
- Two types
  - Executable (Runnable files)
    - Must be called "package main"
    - Must ALWAYS have "func main" as a function
    - GO build makes executable
  - Reusable (Helpers. Reusable logic)
    - "Any other name will make reusable"
    - Makes no executable
    - Files in the same package can freely call functiãons defined in other files. WITHOUT import!

## Run

```bash
go run main.go
go run main.go deck.go
```

## Import

- fmt (format)
- golang.org/pkg (All standart packages)

## Functions

```go
func bobFun() {}

func bobFunction(bobInput string) string {
  return "Hi there"
}

//Receiver function
func (d deck) print() {}
//Any variable of type "deck" now gets access to the print method
```

## Basic types

- bool
- string
- int
- float64

## Vars

```go
var bob string = "Hi I'm bob string"
bob := "bobstring" //(Creating a new var)
bob = "bobstring"  //(Reassigning an existing var)
```

- Can do this, create and assign separatly:

```go
var bob string
bob = "bob"
//This initial creation can be done outside of function scope
//assigning a variable outside of function scope it NOT possible
```

## Data Structures

- Array
  - Fixed length of records

- Slice (0 indexed)
  - Can grow and shrink

```go
bs := make([]byte, 9999)

cards := []string{"bob", "bob2"}
//Returns a new slice of cards with a new element
append(cards, "Six of Spades")


//Sub section
//cards[startIndexIncluding : upToButNotIncluding]
subsection := cards[0:3]
//you can optionally leave out either number e.g. cards[:3] or cards [10:]

//swap positions
cards[i], cards[newPosition] = cards[newPosition], cards[i]
```

- Map
  
```go
//initialize
colours := map[keyType]valueType{}
```

```go
colours := make(map[string]string)

//modification
colours["blue"] = "#0000ff0"
//delete
delete(colours, "blue")
```

## Iterate

```go
for i, card := range cards {
  fmt.Println(i, card)
}
```

## Types

- type bob []string
  - creates a type called bob, used to extend functionality of, in this case, a slice
  - See functions - receiver functions as to how to add functionality

## Testing

- file must end witn "_test.go"
- to run "go test"
- not many testing frameworks available
- test function must start with capital T Test e.g. TestNewDeck

## Struct

```go
type person struct {
 fistName string
 lastName string
}
```

```go
 bob := person{"bob", "bobby"}
```

```go
 bob := person{firstName: "bob",lastName: "bobby"}
 ```

## Zero values

- string ""
- int 0
- float 0
- bool false

## Pointers

- *variable (get the value of the variable)
- &variable (get the address of a variable)
- *type (Used in function, it means it requires a memory address that has a "type" value)
- Gotchas
  - slice[0] = "bob" will update, this is because a slice is a data structure that has a pointer
            to an array, so if a copy is made both copies will have a pointer to the same array
  - this is called a referance type

## Interfaces

```go
type chatBot interface {
 getGreeting() string
 }
 ```

- For anything in the package that has a func callefd getGreeting() that returns a string, then you are also an honorary chatBot

```go
func printGreeting(b chatBot) {
 fmt.Println(b.getGreeting())
 }
 ```

- For anything that is of type chatBot you can now call printGreeting  
- Rules
  - Interfaces are not generics! go does not have generics
  - They are implicit (Don't have to tell "subclasses" about the interface)
  - Can be built out of multiple interfaces

```go
type ReadCloser interface {
  Reader
  Closer
}

type Reader interface {
  Read(p []byte) (n int, err error)
}

type Closer interface {
  Close() error
}
```
