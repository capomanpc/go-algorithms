package main

import "fmt"

//スライス
//copy

func main() {
	/*

		sl := []int{100, 200}
		sl2 := sl
		//このようにsl2を定義するとsl2がslを指すようになる
		//つまり100を指す変数はsl[0]とsl2[0]の二つになってしまう(参照型)

		sl2[0] = 1000
		fmt.Println(sl)

		//基本型の例
		var i int = 10
		i2 := i
		i2 = 100
		fmt.Println(i, i2)

	*/

	//copy関数
	//sliceは参照型なのでコピーする際にはコピー関数を使わなければならない
	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)

	fmt.Println(sl2)

	n := copy(sl2, sl)

	//第一引数はコピー先、第二引数はコピー元
	fmt.Println(n, sl2)
	//nはコピーに成功した数

}
