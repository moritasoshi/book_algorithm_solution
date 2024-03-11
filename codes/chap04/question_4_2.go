package main

import "fmt"

var memo = make(map[int]int)

func main() {
	var N int
	fmt.Scanln(&N)
	ans := tribonacci(N)
	fmt.Println(ans)
}

// トリボナッチ数列
// Tn = Tn-1 + Tn-2 + Tn-3
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if memo[n] != 0 {
		return memo[n]
	}
	val := tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
	memo[n] = val
	return val
}
