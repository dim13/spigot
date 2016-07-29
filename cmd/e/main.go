package main

import (
	"flag"
	"fmt"

	"dim13.org/spigot"
)

var N = flag.Int("n", 60, "number of digits")

func main() {
	flag.Parse()
	for i := range spigot.E(*N) {
		fmt.Print(i)
	}
	fmt.Println("")
}
