package main

import "fmt"

//スライス
//可変長引数
//関数の引数が可変長、何個でも引数として関数に入れることができる

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

/*
func 関数名(スライス名 ...スライスの中身の型)返り値の型{処理内容}

*/

func main() {
	fmt.Println(Sum(1, 2, 3))
	//可変長引数はスライスとして扱われる

	sl := []int{1, 2, 3}
	fmt.Println(Sum(sl...))
}
