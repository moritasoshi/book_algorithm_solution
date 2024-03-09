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
	K, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	count := 0
	for X := 0; X <= K; X++ {
		for Y := 0; Y <= K; Y++ {
			Z := N - (X + Y)
			if Z >= 0 && Z <= K {
				count++
			}
		}
	}
	fmt.Println(count)
}
