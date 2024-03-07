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
	sc.Scan()
	W, _ := strconv.Atoi(sc.Text())
	var a []int
	for i := 0; i < N; i++ {
		sc.Scan()
		val, _ := strconv.Atoi(sc.Text())
		a = append(a, val)
	}

	exists := false
	for bit := 0; bit < (1 << N); bit++ {
		sum := 0
		for i := 0; i < N; i++ {
			if bit&(1<<i) != 0 {
				sum += a[i]
			}
		}
		if sum == W {
			exists = true
		}
	}
	fmt.Println(strconv.FormatBool(exists))
}
