// Package spigot implements Spigot algorithm for Pi and E
package spigot

// spigot step
func spigot(a, r, q int, carry <-chan int) <-chan int {
	c := make(chan int, 10)
	go func() {
		for cr := range carry {
			a = 10*a + cr
			c <- r * (a / q)
			a %= q
		}
		close(c)
	}()
	return c
}

// seed initial values (zero)
func seed(n int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			c <- 0
		}
		close(c)
	}()
	return c
}
