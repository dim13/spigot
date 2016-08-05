package spigot

func filter(a, r, q int, carry <-chan int) <-chan int {
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

func zero(n int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			c <- 0
		}
		close(c)
	}()
	return c
}

func Pi2(n int) <-chan int {
	c := zero(n + 1)
	cr := c
	for i := 10*n/3 + 1; i > 0; i-- {
		cr = filter(2, i, 2*i+1, cr)
	}
	return predigit(filter(2, 1, 10, cr))
}

func E2(n int) <-chan int {
	c := zero(n + 1)
	cr := c
	for i := n + 1; i > 0; i-- {
		cr = filter(1, 1, i+1, cr)
	}
	return filter(2, 1, 10, cr)
}
