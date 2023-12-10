package main

import (
	"flag"
	"fmt"
)

//flag
//コマンドラインの構築ができたり、解析ができる

func main() {
	//コマンドラインのオプション処理
	//コマンドラインを処理するサンプル
	//go run main.go -n 20 -m message -x
	//自分でコマンドのオプションを設定できる
	//上の例では-nや20,-mなどのオプションを入力している
	//これらのオプションは下のIntVar関数やStringVar関数で定義されている

	//オプションの値を格納する変数の定義
	var (
		max int
		msg string
		x   bool
	)

	//IntVar 整数のオプション　var...variable（変数）
	flag.IntVar(&max, "n", 32, "処理数の最大値")

	//StringVar 文字列のオプション
	flag.StringVar(&msg, "m", "", "処理のメッセージ")

	//BoolVar bool型のオプション　コマンドラインに与えられたらTrue なければfalse
	flag.BoolVar(&x, "x", false, "拡張オプション")

	//コマンドラインをパース
	flag.Parse()

	fmt.Println("処理数の最大値=", max)
	fmt.Println("処理のメッセージ=", msg)
	fmt.Println("拡張オプション=", x)
}
