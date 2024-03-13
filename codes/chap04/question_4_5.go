package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scanln(&N)
	count := add357(0, N)
	fmt.Println(count)
}

// add 3 or 5 or 7
func add357(n int, max int) int {
	if n > max {
		return 0
	}
	if is357(n) {
		return 1 + add357(n*10+3, max) + add357(n*10+5, max) + add357(n*10+7, max)
	} else {
		return add357(n*10+3, max) + add357(n*10+5, max) + add357(n*10+7, max)
	}
}

func is357(n int) bool {
	if strings.Contains(strconv.Itoa(n), "3") && strings.Contains(strconv.Itoa(n), "5") && strings.Contains(strconv.Itoa(n), "7") {
		return true
	}
	return false
}
