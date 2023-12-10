package main

import "fmt"

//interface & nil

func main() {
	var x interface{}
	fmt.Println(x)

	//interface{}型	あらゆる型と互換性がある型、数字や文字になれる
	//nil	あらゆる値も持っていない特別な初期状態

	x = 1
	fmt.Println(x)

	x = 3.14
	fmt.Println(x)
	x = "A"
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)

	x = 2
	//fmt.Println(x + 3)
	//printlnメソッド内で演算はできないのに注意

}
