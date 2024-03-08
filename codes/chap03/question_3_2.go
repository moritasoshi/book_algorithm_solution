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
	v, _ := strconv.Atoi(sc.Text())

	var a []int
	for i := 0; i < N; i++ {
		sc.Scan()
		val, _ := strconv.Atoi(sc.Text())
		a = append(a, val)
	}

	var count int
	for _, value := range a {
		if value == v {
			count++
		}
	}

	fmt.Println(count)
}
