package spigot

// Pi calculates n digits of Pi
//
//  1. Initialize: Let A = (2, 2, 2, 2,... ,2) be an array of length
//     [10n/3]+1.
//
//  2. Repeat n times:
//
//     Multiply by 10: Multiply each entry of A by 10.
//
//     Put A into regular form: Starting from the right, reduce the
//     ith element of A (corresponding to b-entry (i - 1)/(2i - 1))
//     modulo 2i - 1, to get a quotient q and a remainder r. Leave
//     r in place and carry q(i - 1) one place left. The last integer
//     carried (from the position where i - 1 = 2) may be as large
//     as 19.
//
//  3. Get the next predigit: Reduce the leftmost entry of A (which
//     is at most 109 (= 9 - 10 + 191)) modulo 10. The quotient, q,
//     is the new predigit of π, the remainder staying in place.
//
//  4. Adjust the predigits: If q is neither 9 nor 10, release the
//     held predigits as true digits of π and hold q. If q is 9, add
//     q to the queue of held predigits. If q is 10 then:
//
//     - set the current predigit to 0 and hold it;
//     - increase all other held predigits by 1(9 becomes 0);
//     - release as true digits of π all but the current held predigit.
//
func Pi(n int) string {
	return format(pi(n))
}

func pi(n int) <-chan int {
	c := seed(n + 1)
	for i := 10*n/3 + 1; i > 0; i-- {
		c = spigot(2, i, 2*i+1, c)
	}
	return predigit(spigot(2, 1, 10, c))
}

func predigit(in <-chan int) <-chan int {
	c := make(chan int)
	drain := func(pre []int, i int) {
		for _, v := range pre {
			c <- (v + i) % 10
		}
	}
	go func() {
		var pre []int
		for i := range in {
			switch i {
			case 9:
				pre = append(pre, i)
			case 10:
				drain(pre, 1)
				pre = append(pre[:0], i%10)
			default:
				drain(pre, 0)
				pre = append(pre[:0], i%10)
			}
		}
		drain(pre, 0)
		close(c)
	}()
	return c
}
