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

	var worst int
	var second int
	for i, value := range a {
		if i == 0 {
			worst = value
			continue
		}
		if value < worst {
			second = worst
			worst = value
			continue
		}
		if worst < value && value < second {
			second = value
			continue
		}
	}

	fmt.Println(second)
}
