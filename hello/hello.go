// Simple Hello application used to demonstrate function calls, and parameters.
package main

import (
	"fmt"
)

// Returns a greeting string to `name`. If `name`
// is empty or nil, defaults to `world`.
func greeting(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(greeting("Kirby"))
}
