package main

import "fmt"

//配列型		要素数を変更できない
//スライス型	要素数を変更可能

func main() {
	var arr1 [3]int
	fmt.Println(arr1)
	//配列の定義

	//arrは[3]int型の変数となる

	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	//要素数を[...]とすることで自動で要素数を指定。型名は...にならない

	fmt.Println(arr2[0])
	//index番号を指定している。先頭はarr2[0]
	fmt.Println(arr2[1])
	fmt.Println(arr2[2]) //空文字が出力される
	fmt.Println(arr2[2-1])

	arr2[2] = "C"
	fmt.Println(arr2[2])

	//var arr5 [4]int
	//arr5 = arr1
	//fmt.Println(arr5)

	//要素数が一致していない場合別の型として認識されるので配列の代入はできない
	fmt.Println(len(arr1))

}
