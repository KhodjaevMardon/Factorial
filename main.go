package main

import "fmt"

func factorial(num int) uint64 {
	var res uint64 = 1
	for i := 2; i < num; i++ {
		res *= uint64(i)
	}
	return res
}

func main() {
	var (
		num int
		res uint64
	)
	fmt.Scan(&num)
	res = factorial(num)
	fmt.Print(num, "! = ", res)
}
