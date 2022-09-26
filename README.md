# Swissknife

A collection of useful functions

## Installation

Install `swissknife` with gopkg:

```bash
go get github.com/ItzAfroBoy/swissknife
```

## Usage

### Lines()

Returns how many lines are in a string

```go
import (
  "github.com/ItzAfroBoy/swissknife"
  "fmt"
)

text := "blah\nblah blah\n\n\nblah blah...\n"
fmt.Println(swissknife.Lines(text)) // 5
```

### Longest()

Returns the length of the longest word or line, as well as the word or line itself.

String:

```go
import (
  "github.com/ItzAfroBoy/swissknife"
  "fmt"
)

text := "apple, pineapple, cheese, milk"
fmt.Println(swissknife.Longest(text, ", ")) // 9, pineapple
```

Line:

```go
import (
  "github.com/ItzAfroBoy/swissknife"
  "fmt"
)

text := "blah\nblah blah\n\n\nblah blah...\n"
fmt.Println(swissknife.Longest(text, "\n")) // 5, blah blah...
```

### Prompt()

Return input from the terminal

```go
import (
  "github.com/ItzAfroBoy/swissknife"
  "fmt"
)

input := swissknife.Prompt("What is your name?") // "What is your name?: " <- Brett
fmt.Println(input) // Brett
```

## License

[GNU GPLv3](https://github.com/ItzAfroBoy/swissknife/blob/main/LICENSE)
