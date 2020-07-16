package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
			k++
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
	pText := []string{}
	for i := 0; i < n; i++ {
		//fmt.Printf("%d ", p[i])
		number := p[i]
		text := strconv.Itoa(number)
		pText = append(pText, text)
	}

	result := strings.Join(pText, ", ")
	fmt.Println(result)
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
	fText := []string{}
	for i := 0; i < n; i++ {
		//fmt.Printf("%v ", f())
		text := strconv.Itoa(f())
		fText = append(fText, text)
	}
	result := strings.Join(fText, ", ")
	fmt.Println(result)
}
