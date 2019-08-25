// Package iteration - sample code for iterations with `for`
package iteration

// Repeat - returns a string where `character` is repeated `n` times
func Repeat(character string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += character
	}
	return result
}
