package main

import (
	"fmt"
	"os"
)

//defer

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

//deferを使うことでTestDefer関数の終了時にdeferの行が実行される
//defer...先送りする、延期する

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

/*
defer文はスタック形式(先入れ後出し)で積まれていく。
1の上に2が積まれて2の上に3が積まれる
実行時は上から実行されるので3,2,1の順で表示される
*/

func main() {
	TestDefer()

	/*

		defer func() {
			fmt.Println("1")
			fmt.Println("2")
			fmt.Println("3")

		}()

	*/

	//無名関数を定義することで複数の処理をdeferで実行できる

	RunDefer()

	//osパッケージのcreate関数でtest.txtファイルを作成

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("HELLO"))
}
