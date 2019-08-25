package variables

import "testing"
import "fmt"

func TestAddr(t *testing.T) {
	expected := 4
	sum := Add(2,2)

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}	
}

// Adds two integers and returns their sum
func ExampleAdd() {
	sum := Add(4, 7)
	fmt.Println(sum)
	// Output: 11
}