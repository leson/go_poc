package greetings

import (
	"fmt"

	"github.com/huandu/xstrings"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	fmt.Sprintln(xstrings.WordCount(message))
	return message
}
