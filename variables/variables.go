// Package variables - ample code to explore Go types and variables
package variables

// Add - adds two integers and returns their sum
func Add(x, y int) int {
	return x + y
}

// AddConv - adds two 'ints' and returns their sum as an 'int32'
func AddConv(x, y int) int32 {
	return int32(x + y)
}
