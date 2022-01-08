package main

import (
	. "fmt"

	"github.com/leson/go_poc/greetings"

	"rsc.io/quote"
)

func main() {
	Println("hello go world!!!")
	Println(quote.Go())
	name := "leson"
	Println(greetings.Hello(name))
}
