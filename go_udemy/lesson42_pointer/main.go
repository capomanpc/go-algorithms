package main

import "fmt"

//ポインタ

func Double(i int) {
	i = i * 2
}

func main() {
	var n int = 100
	fmt.Println(n)

	fmt.Println(&n) //adress表示

	Double(n)

	//main関数の外でnを二倍してもmain関数の中にあるnは二倍されない
	//Double関数の処理はmain関数のnをコピーしているためmain関数のnを直接操作する必要がある

	fmt.Println(n)

	var p *int = &n
	//*int型のpというポインタの宣言
	fmt.Println(p)
	//pの中身はアドレス
	fmt.Println(*p)
	//ポインタ変数に*をつけることでpが指すアドレスに入っている数字等を表示できる

}
