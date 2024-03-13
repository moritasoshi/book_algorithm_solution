package main

import "fmt"

func solve(i int, w int, a []int) bool {
	fmt.Println(i, w)
	// ベースケース
	if i == 0 {
		if w == 0 {
			return true
		} else {
			return false
		}
	}

	// a[i-1] を選ばない場合
	if solve(i-1, w, a) {
		return true
	}

	// a[i-1] を選ぶ場合
	if solve(i-1, w-a[i-1], a) {
		return true
	}

	// どちらも false の場合は false
	return false
}

func main() {
	// 入力
	var N, W int
	//fmt.Scan(&N, &W)
	N, W = 4, 14
	a := make([]int, N)
	for i := 0; i < N; i++ {
		//fmt.Scan(&a[i])
	}
	a = []int{3, 2, 6, 5}

	// 再帰的に解く
	if solve(N, W, a) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
