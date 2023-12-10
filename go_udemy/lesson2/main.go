package main

import "fmt"

//変数

//i5:=500
//暗黙的な定義は関数の外で定義された場合、メイン関数内で表示できない
var i5 int = 500

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)

}

func main() {
	//	明示的な定義
	//var 変数名　型　=　値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	//暗黙的な定義
	//変数名　:=値

	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	//i4 := 500
	//fmt.Println(i4)

	//i4 = "Hello"
	//fmt.Println(i4)

	//暗黙的な定義は後から型を変更できない

	fmt.Println(i5)

	//関数外は明示的な定義
	//関数内は暗黙的な定義
	//Goは静的型付け言語と呼ばれ、型指定をあらかじめ行うことでエラーを減らす

	outer()

	//fmt.Println(s4)

	var s5 string = "Not Use"

	//goでは定義された変数は必ず使う必要がある
	//無駄な変数があるとエラーになる

	fmt.Println(s5)

}
