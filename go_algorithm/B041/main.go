package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	s := []byte(strconv.Itoa(n))
	l := len(s)

	acc := 0
	for i := l; i > 0; i-- {
		if s[i-1] == '1' {
			acc += (1 << (l - i))
		}
	}

	fmt.Println(acc)
}
