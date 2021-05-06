package main

import (
	"fmt"
	"time"
)

func main() {
	delay := 100*time.Millisecond
	go spinner(delay)
	const n = 41
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d", n, fibN)
}

func spinner(delay time.Duration)  {
	for  {
		for _, v := range `-\|/`{
			fmt.Printf("\r%c", v)
			time.Sleep(delay)
		}
	}
}

func fib(n int) int {
	if n < 2{
		return n
	}
	return fib(n - 1) + fib(n - 2)
}