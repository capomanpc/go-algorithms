package main

import (
	"fmt"
	. "fmt" //ピリオドを付けることでfmtパッケージの関数を使う際にパッケージ名を省略して関数を使うことができる
	f "fmt" //fmtパッケージに任意の名前fを付けることができる
	"go_udemy/lesson53_scope/foo"
)

//スコープ

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	//var b string = s
	b = s
	{
		b := "BBBB"
		fmt.Println(b)
		//関数内で同じ変数を使用したい場合{}でブロックを作れば湧けることができる
	}
	fmt.Println(b)
	return b
}

//関数のシグニチャで返り値の変数が指定されている場合に、関数内で同じ変数を定義すると再定義になってしまいエラーとなる

func main() {
	Println(foo.Max)
	//同一フォルダ内にあるfooパッケージのMaxを参照
	//一文字目が大文字であれば他パッケージからでも参照が可能

	//fmt.Println(foo.min)
	//上記コードは一文字目が小文字なので他パッケージから参照できない

	f.Println(foo.ReturnMin())
	//fooパッケージのReturnMin関数を実行することで一文字目が小文字でもminの値を得ることができる
	//ReturnMin関数も一文字目が大文字なので他パッケージ空でも呼び出すことができる

	//大文字からスタート -> パブリックな変数、関数
	//小文字からスタート -> プライベートな変数、関数

	fmt.Println(appName())

	fmt.Println(Do("AAA"))

}
