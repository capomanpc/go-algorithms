package main

import "fmt"

func main() {
	var n, q, l, r int
	fmt.Scan(&n)
	a := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
		a[i] += a[i-1]
	}
	fmt.Scan(&q)

	for i := 1; i <= q; i++ {
		fmt.Scan(&l, &r)
		winCount := a[r] - a[l-1]
		if winCount == 0 {
			fmt.Println("draw")
		} else if winCount > 0 {
			fmt.Println("win")
		} else {
			fmt.Println("lose")
		}
	}
}
