package main

import (
	"flag"

	"dim13.org/spigot"
)

func main() {
	N := flag.Int("n", 60, "number of digits")
	flag.Parse()
	spigot.Print(spigot.Pi(*N))
}
