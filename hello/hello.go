// Simple Hello application used to demonstrate function calls, and parameters.
package main

import (
	"fmt"
)

// Returns a greeting string to `name`. If `name`
// is empty defaults to `World`.
func greeting(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello " + name
}

func main() {
	fmt.Println(greeting("Kirby"))
}
