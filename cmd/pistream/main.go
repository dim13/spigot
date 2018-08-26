package main

import (
	"fmt"

	"github.com/dim13/spigot"
)

func main() {
	c := make(chan int64)
	go spigot.PiStream(c)
	for i := 1; ; i++ {
		fmt.Print(<-c)
		if i%10 == 0 {
			fmt.Printf("\t:%d\n", i)
		}
	}
}
