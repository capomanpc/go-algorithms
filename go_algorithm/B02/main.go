package main

import "fmt"

func main() {
	var A, B int
	var canDivide bool

	fmt.Scan(&A, &B)

	for i := A; i <= B; i++ {
		if 100%i == 0 {
			canDivide = true
		}
	}

	if canDivide {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
