package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var n, j int
	fmt.Scan(&n)

	nStr := []byte(strconv.Itoa(n))
	length := len(nStr)

	//integer to ASCII(整数から文字列)したものをbyte配列に代入

	ans := 0.0

	for i := length - 1; i >= 0; i-- {
		if nStr[i] == '1' {
			ans += math.Pow(2, float64(j))
		}
		j++
	}
	fmt.Println(ans)
}
