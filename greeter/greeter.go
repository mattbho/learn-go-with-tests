package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	// "os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprint(writer, "Hello, ", name)
}

func MyGreeterHandler (w http.ResponseWriter, r *http.Request){
	Greet(w, "Matt")
}

func main() {
	// Greet(os.Stdout, "Matt") // os.Stdout implements io.Writer
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
