package main

import "fmt"

func fibonacci() func() int {
	i := 0
	a := 0
	b := 1

	return func() int {
		i = a
		a = a + b
		b = i
		return i
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
