package main

import "fmt"

//スライス
//append make len cap

func main() {
	sl := []int{100, 200}
	fmt.Println(sl)

	//appendアペンド拡張するための関数
	//スライスと配列の違いは拡張性、配列は要素数指定する必要があるが、スライスは指定する必要がなく制限がない
	sl = append(sl, 300)
	fmt.Println(sl)

	sl = append(sl, 400, 500, 600)
	fmt.Println(sl)

	//make関数スライス生成
	sl2 := make([]int, 5)
	fmt.Println(sl2)

	//len関数
	fmt.Println(len(sl2))

	//cap関数
	fmt.Println(cap(sl2))

	//5が最初の要素数、10がキャパシティーで最大で10の要素を格納できる領域を確保
	sl3 := make([]int, 5, 10)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))

	//make([]int,5,10)でキャパシティーを10個確保した後にappendで10を超える要素を追加するとキャパシティーは二倍になってしまう
	//過剰なメモリ確保は実行速度が落ちたりする
	sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))

}
