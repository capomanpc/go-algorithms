package main

import "fmt"

func main() {
	var n, k, count int
	fmt.Scan(&n, &k)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			z := k - i - j
			if z >= 1 && z <= n {
				count++
			}
		}
	}
	fmt.Println(count)
}
