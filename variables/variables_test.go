package variables

import "testing"
import "fmt"

func TestAddr(t *testing.T) {
	expected := 4
	sum := Add(2, 2)

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestConvInt32(t *testing.T) {
	want := "int32"
	got := fmt.Sprintf("%T", AddConv(16, 21))
	if want != got {
		t.Errorf("want % q, got %q", want, got)
	}
}

// Adds two integers and returns their sum
func ExampleAdd() {
	sum := Add(4, 7)
	fmt.Println(sum)
	// Output: 11
}

// Adds two integers (int) and returns their sum as an (int32)
func ExampleAddConv() {
	var sum int32 = AddConv(16, 30)
	fmt.Println(sum)
	// Output: 46
}
