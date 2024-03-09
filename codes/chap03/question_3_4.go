package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	var a []int
	for i := 0; i < N; i++ {
		sc.Scan()
		val, _ := strconv.Atoi(sc.Text())
		a = append(a, val)
	}

	max := a[0]
	min := a[0]

	for i := 0; i < N; i++ {
		// 最大値
		if a[i] > max {
			max = a[i]
		}
		// 最小値
		if a[i] < min {
			min = a[i]
		}
	}

	diff := max - min
	fmt.Println(diff)
}
