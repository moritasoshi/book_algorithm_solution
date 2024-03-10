package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scanln(&S)
	N := len(S)

	// N-1 箇所の隙間に「+」を入れるかどうかをすべて試す
	res := int64(0)
	for bit := 0; bit < (1 << (N - 1)); bit++ {
		// tmp：「+」と「+」との間の値を表す変数
		tmp := int64(0)
		for i := 0; i < N-1; i++ {
			tmp *= 10
			tmp += int64(S[i] - '0')
			// 「+」を挿入するとき、答えに tmp を加算して、tmp を 0 に初期化
			if (bit & (1 << i)) != 0 {
				res += tmp
				tmp = 0
			}
		}
		tmp *= 10
		tmp += int64(S[N-1] - '0')
		res += tmp
	}

	fmt.Println(res)
}
