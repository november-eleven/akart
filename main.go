package main

import (
	"fmt"
	"os"
)

// Launch akart engine and parse application's arguments as keyword to auto-complete.
func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s keyword...\n", os.Args[0])
		os.Exit(255)
		return
	}

	e := NewEngine()

	for _, k := range os.Args[1:] {
		v := e.Find(k)
		fmt.Printf("%s: %v\n", k, v)
	}

}
