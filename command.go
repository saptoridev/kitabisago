package main

import "fmt"

type Command interface {
	Execute()
}

type SumXY struct {
	X int
	Y int
}

type MultiplyXY struct {
	X int
	Y int
}

type PrimeNumber struct {
	N int
}

type FibonacciNumber struct {
	N int
}

func (i SumXY) Execute() {
	fmt.Printf("Input: %d, %d\n", i.X, i.Y)
	fmt.Printf("Output: %v\n", sumXY(i.X, i.Y))
}

func (i MultiplyXY) Execute() {
	fmt.Printf("Input: %d, %d\n", i.X, i.Y)
	fmt.Printf("Output: %v\n", multiplyXY(i.X, i.Y))
}

func (i PrimeNumber) Execute() {
	fmt.Printf("Input: %d\n", i.N)
	fmt.Printf("Output: ")
	primeSequence(i.N)
}

func (i FibonacciNumber) Execute() {
	fmt.Printf("Input: %d\n", i.N)
	fmt.Printf("Output: ")
	fibonacciSequence(i.N)
}

type Invoker struct {
	X int
	Y int
	N int
}

func (i Invoker) ExecuteCommands() {
	var c Command

	c = SumXY{X: i.X, Y: i.Y}
	fmt.Printf("Sum X & Y \n")
	c.Execute()

	fmt.Printf("\n")

	c = MultiplyXY{X: i.X, Y: i.Y}
	fmt.Printf("Multiply X & Y \n")
	c.Execute()

	fmt.Printf("\n")

	c = PrimeNumber{N: i.N}
	fmt.Printf("First N prime number \n")
	c.Execute()

	fmt.Printf("\n")

	c = FibonacciNumber{N: i.N}
	fmt.Printf("First N Fibonacci sequence, \n")
	c.Execute()

}
