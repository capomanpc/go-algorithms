package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//ファイル操作
	//Read
	f, err := os.Open("foo.txt")

	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	bs := make([]byte, 128)

	n, err := f.Read(bs)
	//ファイルfを読み込んでbsスライスに格納
	//nは読み込んだバイト数を返す

	fmt.Println(n)
	fmt.Println(string(bs))
	//byte[]スライスに入った値を文字列に変換して表示

	bs2 := make([]byte, 128)

	nn, err := f.ReadAt(bs2, 10)
	//第二引数はオフセット位置を指定(10byte目から読み取り)

	fmt.Println(nn)
	fmt.Println(string(bs2))
}
