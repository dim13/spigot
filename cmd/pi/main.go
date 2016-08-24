package main

import (
	"flag"
	"fmt"

	"dim13.org/spigot"
)

func main() {
	N := flag.Int("n", 60, "number of digits")
	flag.Parse()
	for v := range spigot.Pi(*N) {
		fmt.Print(v)
	}
	fmt.Println("")
}
