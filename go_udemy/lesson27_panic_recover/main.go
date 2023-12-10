package main

import "fmt"

//panic&recover
//プログラムを強制終了する強力な構文なためあまり使わないことが望ましい

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	//defer文で無名関数を定義
	//recover()関数はパニック発生時にパニック情報を取得する関数
	//x != nil のときすなわちxにパニック情報が代入されたときにパニック情報を表示する

	//無名関数の最後に()をつけることで即座に実行出来る

	panic("runtime error")
	fmt.Println("Start")

	//ここでpanic(ランタイムエラー)を起こしている
	//panic関数は強制的にランタイムエラーを起こす関数

	//defer文により遅延実行をすることでパニック関数によりランタイムエラーを起こしてもプログラムが止まらずに適切な処置ができる

}
