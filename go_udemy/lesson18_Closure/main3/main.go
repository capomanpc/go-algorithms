package main

import "fmt"

func Counter() int {
	var count int = 0

	a := func() int {
		count++
		return count
	}
	b := a()

	return b
}

func main() {

	a := Counter()
	fmt.Println(a)
	fmt.Println(a)

}

/*
このコードは関数Couterの結果の整数値をaに代入して表示している
aが使われる度に関数Counterが実行されている
そのため実行される度に変数countが0で初期化されてしまうためクロージャになっていない
*/
