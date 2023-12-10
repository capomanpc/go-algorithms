package main

import "fmt"

//関数
//クロージャー

func Later() func(string) string {
	//returnで無名関数をその場で定義してその無名関数をそのまま返す関数
	var store string
	//空文字をstoreに代入
	return func(next string) string {
		//nextという引数を受け取る
		s := store
		store = next
		//sにsroreをコピーしてから引数のnextを代入
		return s
	}
}

/*
Later関数は引数がなく、返り値がfunc(string)string型の関数
つまり返り値はstring型を引数に受け取ってstring型の返り値を返す関数
*/

func main() {
	f := Later()

	/*
		Later()の結果がfに代入されるため、fのデータ型はLaterの返り値と一致する
		fmt.Println(Later()("Hello"))としてもよい。こうすることでnextに"Hello"を代入できる
		f := Later()とすることで、fはLater関数の結果が代入された変数となる。
		つまり結果だけになり扱いやすくなる。
	*/

	fmt.Println(f("Hello"))
	fmt.Println(f("My"))
	fmt.Println(f("Name"))
	fmt.Println(f("is"))
	fmt.Println(f("Golang"))

}

/*
このコードでは、関数Laterがクロージャを返している。
つまりこの場合のクロージャはLater内で定義された無名関数のこと。

クロージャは、自身が定義されたスコープ内の変数にアクセスできる関数のことである。
この場合は自身が定義されたスコープ内の変数というのはstoreのこと。

f := Later()がない場合を考える。
fを使わない場合fmt.Println(Later()("Hello"))と書き換えることができる。
このように書き換えると出力結果は全て空文字になってしまう(main2.goを参照)。
空文字になる理由はその場その場で関数を呼び出しているため、呼び出す度に初期化されてしまうからである

f := Later()がある場合を考える。
fがあると、実質的にreturnにある無名関数だけが実行される
そのとき無名関数は前に使った変数の中身は破棄せずに持ち続けているので





*/
