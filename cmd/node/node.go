package main

import (
	"fmt"
	"time"

	"github.com/freitzzz/go-zeromq-auto-discover/cmd/node/methods"
)

func main() {
	methods := map[string]func(){
		"1. Hardcoded endpoints":     methods.Hardcoded,
		"2. Configuration endpoints": methods.Configuration,
		"3. Beacon discovery":        methods.Beacon,
	}

	for n, m := range methods {
		fmt.Printf("Running method: %s\n", n)
		t := time.Now()
		m()
		d := time.Since(t)

		fmt.Printf("Took %d ms to run\n", d.Milliseconds())
	}
}
