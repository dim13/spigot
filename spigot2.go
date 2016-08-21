package spigot

func spigot(a, r, q int, carry <-chan int) <-chan int {
	c := make(chan int, 100)
	go func() {
		defer close(c)
		for cr := range carry {
			a = 10*a + cr
			c <- r * (a / q)
			a %= q
		}
	}()
	return c
}

func zero(n int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < n; i++ {
			c <- 0
		}
	}()
	return c
}

// Pi2 calculates n digits of Pi concurently
func Pi2(n int) <-chan int {
	c := zero(n + 1)
	cr := c
	for i := 10*n/3 + 1; i > 0; i-- {
		cr = spigot(2, i, 2*i+1, cr)
	}
	return predigit(spigot(2, 1, 10, cr))
}

// E2 calculates n digits of E concurently
func E2(n int) <-chan int {
	c := zero(n + 1)
	cr := c
	for i := n + 1; i > 0; i-- {
		cr = spigot(1, 1, i+1, cr)
	}
	return spigot(2, 1, 10, cr)
}
