package main

import (
	"fmt"
	"os"
)

//os
//エラー発生時の即時終了や即時終了をしたいときに使用する強制終了のための関数

func main() {
	//Exit
	//main関数の処理が終了するとプログラムが終了するがos.Exit()を使うことで任意の場所で終了できる
	//引数を指定することでexsit status(終了ステータス)を指定できる

	//0:プログラムが問題なく終了したとき0を返して終了する
	//1:エラーが発生した場合に1(非0)を返して終了する
	//自分で0か1かを指定してエラーステータスをプログラムする

	/*
		os.Exit(1)
		fmt.Println("Start")
	*/

	defer func() {
		fmt.Println("defer")
	}()
	os.Exit(0)

	//defer文は実行されない

}
