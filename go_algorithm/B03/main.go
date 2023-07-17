package main

import "fmt"

func main() {
	var N int
	var A [100]int

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	for i := 0; i < N-2; i++ {
		for j := i + 1; j < N-1; j++ {
			for k := j + 1; k < N; k++ {
				if A[i]+A[j]+A[k] == 1000 {
					fmt.Println("Yes")
					return
				}
			}
		}
	}
	fmt.Println("No")
}
