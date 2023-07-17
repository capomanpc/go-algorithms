package main

import "fmt"

func main() {
	var N, X, A int
	var result string = "No"

	fmt.Scan(&N, &X)

	for i := 1; i <= N; i++ {
		fmt.Scan(&A)
		if A == X {
			result = "Yes"
		}
	}
	fmt.Println(result)
}
