package main

import (
	"fmt"
	"os"
)

func main() {
	//Args = Arguments(引数)
	//コマンドライン引数を格納するスライスがos.Args[]
	//コマンドライン引数とはコマンドで入力された引数のこと

	//コマンドで入力した引数を自動でスライスに格納

	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])

	//os.Argsの要素数を表示
	fmt.Printf("length=%d\n", len(os.Args))
	//osの中身を全て表示
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
	/*

		0 getargs.exe
		1 123
		2 456
		3 789
	*/

}
