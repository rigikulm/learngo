package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Method: %v\n", r.Method)
	fmt.Printf("URL: %s\n", r.URL)
	fmt.Printf("Headers: %s\n", r.Header)
	Greet(w, "web world!")
}

func main() {
	http.ListenAndServe(":9090", http.HandlerFunc(MyGreetHandler))
	//Greet(os.Stdout, "Elodie")
}
