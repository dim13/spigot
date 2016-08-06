package main

import (
	"flag"

	"dim13.org/spigot"
)

var (
	N = flag.Int("n", 60, "number of digits")
	C = flag.Bool("c", false, "concurent")
)

func main() {
	flag.Parse()
	if *C {
		spigot.Print(spigot.E2(*N))
	} else {
		spigot.Print(spigot.E(*N))
	}
}
