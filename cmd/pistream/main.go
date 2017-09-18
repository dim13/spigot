package main

import (
	"fmt"

	"dim13.org/spigot"
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
