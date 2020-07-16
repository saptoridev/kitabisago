package main

import "fmt"

func sumXY(x, y int) int {
	return x + y
}

func multiplyXY(x, y int) int {
	return x * y
}

func primes(n int) []int {
	var retval = make([]int, n)
	var k = 0

	for x := 2; ; x++ {
		isPrime := true
		for y := 2; y < x; y++ {
			if x%y == 0 {
				isPrime = false
				break
			}
		}
		if isPrime && k < n {
			retval[k] = x
			k += 1
			continue
		}
		if k >= n {
			break
		}
	}

	return retval
}

func primeSequence(n int) {
	p := primes(n)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", p[i])
	}
}

func fibonacciClosure() func() int {
	first, second := 0, 1
	return func() int {
		retval := first
		first, second = second, first+second
		return retval
	}
}

func fibonacciSequence(n int) {
	f := fibonacciClosure()
	for i := 0; i < n; i++ {
		fmt.Printf("%v ", f())
	}
}

func main() {
	x := 1
	y := 2

	fmt.Printf("sum %d + %d  = %v\n", x, y, sumXY(x, y))
	fmt.Printf("multiply %d * %d = %v\n", x, y, multiplyXY(x, y))
	primeSequence(4)
	fmt.Println()
	fibonacciSequence(4)

}
