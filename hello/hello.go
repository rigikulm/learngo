// Simple Hello application used to demonstrate function calls, and parameters.
package main

import (
	"fmt"
)

// Returns a greeting string
func greeting() string {
	return "Hello World"
}

func main() {
	fmt.Println(greeting())
}
