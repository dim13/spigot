package main

import (
	"flag"
	"fmt"

	"github.com/dim13/spigot"
)

func main() {
	N := flag.Int("n", 60, "number of digits")
	flag.Parse()
	fmt.Println(spigot.E(*N))
}
