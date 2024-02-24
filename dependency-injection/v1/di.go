package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Greet prints a greeting message to the writer with the provided name.
//
// writer: an io.Writer to write the greeting message to
// name: a string representing the name of the person to greet
// Return type: void
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// http.ResponseWriter also implement io.Writer interface
func greetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Felix")
}

func main() {
	Greet(os.Stdout, "Felix")

	log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(greetHandler)))
}
