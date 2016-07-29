package spigot

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

func Pi(n int) <-chan int {
	c := make(chan int)
	go func(n int) {
		l := 10*n/3 + 1
		a := make([]int, l)
		b := make([]int, l)
		for i := 0; i < l; i++ {
			a[i] = 2
			b[i] = 2*i + 1
		}
		for k := 0; k < n; k++ {
			for i := range a {
				a[i] *= 10
			}
			for i := len(a) - 1; i > 0; i-- {
				a[i-1] += i * (a[i] / b[i])
				a[i] %= b[i]
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
   1. Initialize: Let the ﬁrst digit be 2 and initialize an array
      A of length n + 1 to (1, 1, 1, . . . , 1).

   2. Repeat n − 1 times:

      Multiply by 10: Multiply each entry of A by 10.

      Take the fractional part: Starting from the right, reduce the
      ith entry of A modulo i + 1, carrying the quotient one place
      left.

      Output the next digit: The ﬁnal quotient is the next digit of e.
*/

func E(n int) <-chan int {
	c := make(chan int)
	go func(n int) {
		l := n + 1
		a := make([]int, l)
		b := make([]int, l)
		for i := 0; i < l; i++ {
			a[i] = 1
			b[i] = i + 1
		}
		a[0] = 2
		for k := 0; k < n; k++ {
			for i := range a {
				a[i] *= 10
			}
			for i := len(a) - 1; i > 0; i-- {
				a[i-1] += a[i] / b[i]
				a[i] %= b[i]
			}
			c <- a[0] / 10
			a[0] %= 10
		}
		close(c)
	}(n + 1)
	return predigit(c)
}
