package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	res_prev := 0
	res := 1
	return func(x int) int {
		if x == 0 {
			return res_prev
		} else if x == 1 {
			return res
		} else {
			temp := res
			res = res + res_prev
			res_prev = temp
		}
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
