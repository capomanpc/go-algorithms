package main

import (
	"fmt"
)

func main() {
	var n int
	var ans string

	fmt.Scan(&n)

	for i := 10; i > 0; i-- {
		if n%2 == 1 {
			ans = "1" + ans
		} else {
			ans = "0" + ans
		}
		n = n / 2
	}
	fmt.Println(ans)
}
