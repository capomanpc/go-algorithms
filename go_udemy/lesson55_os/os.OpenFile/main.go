package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//ファイル操作
	//OpenFile
	//O_RDONLY 読み込み専用
	//O_WRONLY 書き込み専用
	//O_RDWR 読み書き専用
	//O_APPEND ファイル末尾に追記
	//O_CREATE ファイルが無ければ作成
	//O_TRUNC 可能であればファイルの内容をオープン時に空にする
	//O_と付いているのはOpenFileと関係しているため

	f, err := os.OpenFile("foo.txt", os.O_RDONLY, 0666)
	//f,err := os.OpenFile(ファイル名,フラグ,パーミッション)
	//fはファイルへのハンドラ、フラグとは特定のオプションや設定を表すもの
	//0666　linuxと同様のファイルの管理者設定

	//第三引数の0666はファイル作成時(O_CREATE)のみに使われる
	//今回の場合は第三引数の0666は使われない
	//しかし第三引数を書かないとエラーになってしまうので慣例として0666を引数として渡す

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))
}

/*
	os.File型の定義
	構造体で新しく型を定義している

	type File struct {
		*file
	}

*/