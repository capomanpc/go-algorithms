package main

import (
	"fmt"
)

//型変換

func main() {
	/*
		var i int = 1
		fl64 := float64(i)
		fmt.Println(fl64)
		fmt.Printf("i = %T\n", i)
		fmt.Printf("fl64 = %T\n", fl64)

		//数値型の相互変換

		i2 := int(fl64)
		fmt.Printf("i2 = %T\n", i2)

		fl := 10.5
		i3 := int(fl) //小数点以下の切り捨て
		fmt.Printf("i3=%T\n", i3)
		fmt.Println(i3)
	*/

	//文字列型から数値型への変換

	var s string = "100"
	fmt.Printf("s=%T\n", s)

	/*
		i, err := strconv.Atoi("A")
		//エラーハンドリング
		if err != nil {
			fmt.Println("error")
		}
	*/

	/*ストリングコンバートパッケージの
	 エートゥーアイ関数(メソッド)

	 string型の変数をint型に変換
	 i, err := strconv.Atoi(s)

	 i		正しく変換できたときにiに代入
	 err	二つ目の返り値はerror型の変数になる(返り値はnil)

	 _	正しく変換できなかったときに返り値を破棄できる
		プログラム中に使われていない変数があると正しく動作
		しないため変数を破棄する

	*/

	/*

		fmt.Println(i)
		fmt.Printf("i=%T\n", i)

		//数値型から文字列型への変換

		var i2 int = 200
		s2 := strconv.Itoa(i2)
		fmt.Println(s2)
		fmt.Printf("s2=%T\n", s2)

	*/

	//文字列とbyte配列の変換
	//byte型	1バイトの整数を格納するデータ型で文字に変換することができる

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	//byte型→文字列型
	h2 := string(b)
	fmt.Println(h2)

}
