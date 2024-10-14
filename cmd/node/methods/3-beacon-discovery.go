package methods

import (
	"log"

	"github.com/zeromq/gyre"
)

func Beacon() {
	g, err := gyre.New()
	if err != nil {
		log.Fatal(err)
	}

	defer g.Stop()

	err = g.Start()
	if err != nil {
		log.Fatal(err)
	}

	t := make(chan (string))
	go func() {
		for e := range g.Events() {
			if e.Type() == gyre.EventEnter {
				t <- "!"
			}
		}
	}()

	<-t
}
