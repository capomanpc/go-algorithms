package main

import (
	"fmt"
	"strconv"
)

//if
//条件分岐
//エラーハンドリング

func main() {
	var s string = "a"

	i, err := strconv.Atoi(s)

	//errを明示的に宣言すると以下のようになる
	//var err error

	//_	はエラーが出た場合に破棄するという意味
	//errはinterface{}型の変数名なのでERRでもerrorでも自由な名づけができる

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i=%T", i)

}

/*
strconvパッケージのAtoi関数
string convertion(文字列変換)
ASCII to integer の略がAtoi
*/
