package main

import "fmt"

//標準入力
//standard input

func main() {
	var N int
	fmt.Scan(&N)
	//空白・Tab・改行までの入力を読み取る

	fmt.Scanf("%d", &N)
	//指定したフォーマットに従って変数に値を代入

	fmt.Scanln(&N)
	//改行までの入力を読み取る

	var A, B int

	fmt.Scan(&A, &B)
	//スペースを入力することで2つ定義できる

	fmt.Scanf("%d,%d", &A, &B)
	//フォーマット指定をする場合カンマで二つ入力

}
