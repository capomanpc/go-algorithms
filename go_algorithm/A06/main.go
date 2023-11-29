package main

import "fmt"

func main() {
	var n, q, l, r int

	fmt.Scan(&n, &q)

	a := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
		a[i] += a[i-1]
	}

	for i := 1; i <= q; i++ {
		fmt.Scan(&l, &r)
		fmt.Println(a[r] - a[l-1])
	}
}
