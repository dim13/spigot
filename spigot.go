// Package spigot implements Spigot algorithm for Pi and E
package spigot

import "fmt"

/*
   1. Initialize: Let A = (2, 2, 2, 2,... ,2) be an array of length
      [10n/3]+1.

   2. Repeat n times:

      Multiply by 10: Multiply each entry of A by 10.

      Put A into regular form: Starting from the right, reduce the
      ith element of A (corresponding to b-entry (i - 1)/(2i - 1))
      modulo 2i - 1, to get a quotient q and a remainder r. Leave
      r in place and carry q(i - 1) one place left. The last integer
      carried (from the position where i - 1 = 2) may be as large
      as 19.

   3. Get the next predigit: Reduce the leftmost entry of A (which
      is at most 109 (= 9 - 10 + 191)) modulo 10. The quotient, q,
      is the new predigit of π, the remainder staying in place.

   4. Adjust the predigits: If q is neither 9 nor 10, release the
      held predigits as true digits of π and hold q. If q is 9, add
      q to the queue of held predigits. If q is 10 then:

      - set the current predigit to 0 and hold it;
      - increase all other held predigits by 1(9 becomes 0);
      - release as true digits of π all but the current held predigit.

*/

// alloc allocates initial slice of size n
func alloc(n, v int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = v
	}
	return a
}

// Pi returns n digits of Pi
func Pi(n int) <-chan int {
	c := make(chan int)
	go func(n int) {
		a := alloc(10*n/3+1, 2)
		for k := 0; k < n; k++ {
			a[len(a)-1] *= 10
			for i := len(a) - 1; i > 0; i-- {
				q := 2*i + 1
				a[i-1] *= 10
				a[i-1] += i * (a[i] / q)
				a[i] %= q
			}
			c <- a[0] / 10
			a[0] %= 10
		}
		close(c)
	}(n + 1)
	return predigit(c)
}

func predigit(in <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		var pre []int
		for i := range in {
			switch i {
			case 9:
				pre = append(pre, i)
			case 10:
				for _, v := range pre {
					c <- (v + 1) % 10
				}
				pre = []int{0}
			default:
				for _, v := range pre {
					c <- v
				}
				pre = []int{i}
			}
		}
		for _, v := range pre {
			c <- v
		}
		close(c)
	}()
	return c
}

/*
   1. Initialize: Let the first digit be 2 and initialize an array
      A of length n + 1 to (1, 1, 1, . . . , 1).

   2. Repeat n − 1 times:

      Multiply by 10: Multiply each entry of A by 10.

      Take the fractional part: Starting from the right, reduce the
      ith entry of A modulo i + 1, carrying the quotient one place
      left.

      Output the next digit: The final quotient is the next digit of e.
*/

// E returns n digits of E
func E(n int) <-chan int {
	c := make(chan int)
	go func(n int) {
		a := alloc(n+1, 1)
		a[0] = 2
		for k := 0; k < n; k++ {
			a[len(a)-1] *= 10
			for i := len(a) - 1; i > 0; i-- {
				q := i + 1
				a[i-1] *= 10
				a[i-1] += a[i] / q
				a[i] %= q
			}
			c <- a[0] / 10
			a[0] %= 10
		}
		close(c)
	}(n + 1)
	return c
}

// Print digits from channel
func Print(n <-chan int) {
	for v := range n {
		fmt.Print(v)
	}
	fmt.Println("")
}

// Drain values from channel
func Drain(c <-chan int) {
	for range c {
	}
}
