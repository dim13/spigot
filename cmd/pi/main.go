package main

import (
	"fmt"

	"dim13.org/spigot"
)

func main() {
	for i := range spigot.Pi(60) {
		fmt.Print(i)
	}
	fmt.Println("")
}
