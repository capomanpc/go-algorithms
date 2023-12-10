package main

import (
	"log"
	"os"
)

//os

func main() {
	//log.Fatal、強制的にエラーを引き起こしてプログラムを強制終了する関数
	_, err := os.Open("A.txt")
	if err != nil {
		log.Fatalln(err)
		//log.Fatal関数で指定した引数が標準エラー出力される
		//その後内部でos.Exit(1)が呼び出されプログラムは強制終了される
	}
}
