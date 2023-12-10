package main

import "fmt"

//関数
//クロージャー

func Later2() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	fmt.Println(Later2()("Hello"))
	fmt.Println(Later2()("My"))
	fmt.Println(Later2()("Name"))
	fmt.Println(Later2()("is"))
	fmt.Println(Later2()("Golang"))
}

/*
このコードはf:=Later()を消してLater2()("Hello")のように中身の無名関数に直接"Hello"を代入している
これを実行するとPrintlnは全て空文字が出力されてしまう

そうなってしまう理由は毎回毎回Later()を実行しているため
storeの値が毎回毎回0に初期化されてしまっているからである

毎回毎回0に初期化されないためにもf:=Later()を追加してLaterの出力結果の関数をfに保存して使う
そうすることでLater関数を実行してもvar store stringの処理がされないため初期化はされない



*/
