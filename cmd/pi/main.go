package main

import (
	"flag"

	"dim13.org/spigot"
)

var N = flag.Int("n", 60, "number of digits")

func main() {
	flag.Parse()
	spigot.Print(spigot.Pi(*N))
}
