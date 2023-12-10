package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	nStr := []byte(strconv.Itoa(n))
	length := len(nStr)

	ans := 0

	for i := length - 1; i >= 0; i-- {
		if nStr[i] == '1' {
			ans += (1 << (length - i - 1))
		}
	}
	//左シフト演算子を使うと値が二倍になる
	fmt.Println(ans)
}
