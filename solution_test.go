package main

import "testing"

func TestSumXY(t *testing.T) {
	output := sumXY(1, 2)
	valid := 3
	if output != valid {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", output, valid)
	}
}

func TestMultiplyXY(t *testing.T) {
	output := multiplyXY(1, 2)
	valid := 2
	if output != valid {
		t.Errorf("Multipy was incorrect, got: %d, want: %d.", output, valid)
	}
}

func TestPrimes(t *testing.T) {
	output := primes(4)
	valid := [4]int{2, 3, 5, 7}

	if len(output) != len(valid) {
		t.Errorf("Prime sequence was incorrect, got: %v, want: %v.", output, valid)
	}

	for i, v := range output {
		if v != valid[i] {
			t.Errorf("Prime sequence was incorrect, got: %v, want: %v.", output, valid)
		}
	}

}

func TestFibonacciClosure(t *testing.T) {
	n := 4
	valid := [13]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}
	f := fibonacciClosure()

	for i := 0; i < n; i++ {
		output := f()
		if output != valid[i] {
			t.Errorf("Fibonacci sequence was incorrect, got: %v, want: %v.", output, valid[i])
		}
	}

}
