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

	count := 0
out:
	for {
		for i := 0; i < N; i++ {
			// 奇数の場合
			if a[i]%2 != 0 {
				break out
			}
			a[i] = a[i] / 2
		}
		count++
	}
	fmt.Println(count)
}
