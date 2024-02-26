package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func Countdown(writer io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)
		// add 1-second pause
		// we have a dependency on Sleeping which we need to extract
		// so we can then control it in our test
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(writer, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
