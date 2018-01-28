package main

import (
	"flag"
	"fmt"

	"dim13.org/spigot"
)

func main() {
	N := flag.Int("n", 60, "number of digits")
	flag.Parse()
	fmt.Println(spigot.E(*N))
}
