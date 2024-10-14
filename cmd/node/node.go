package main

import (
	"fmt"
	"time"

	"github.com/freitzzz/go-zeromq-auto-discover/cmd/node/methods"
)

func main() {
	methods := map[string]func(){
		"Hardcoded endpoints": methods.Hardcoded,
	}

	for n, m := range methods {
		fmt.Printf("Running method: %s\n", n)
		t := time.Now()
		m()
		d := time.Since(t)

		fmt.Printf("Took %d microseconds to run\n", d.Microseconds())
	}
}
