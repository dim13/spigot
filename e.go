package spigot

// E calculates n digits of E
//
//  1. Initialize: Let the first digit be 2 and initialize an array
//     A of length n + 1 to (1, 1, 1, . . . , 1).
//
//  2. Repeat n − 1 times:
//
//     Multiply by 10: Multiply each entry of A by 10.
//
//     Take the fractional part: Starting from the right, reduce the
//     ith entry of A modulo i + 1, carrying the quotient one place
//     left.
//
//     Output the next digit: The final quotient is the next digit of e.
//
func E(n int) string {
	return format(e(n))
}

func e(n int) <-chan int {
	c := seed(n + 1)
	for i := n + 1; i > 0; i-- {
		c = spigot(1, 1, i+1, c)
	}
	return spigot(2, 1, 10, c)
}
