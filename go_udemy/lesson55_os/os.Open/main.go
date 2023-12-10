package main

import (
	"log"
	"os"
)

//ファイルを読み込み専用モードでopenしてos.File型のファイルへのハンドラ変数fを取得

func main() {
	//ファイル操作
	//os.Open()

	f, err := os.Open("test.txt")
	//ファイル名, エラー := os.Open()

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()
}

/*
	ファイルを開くときには読み込み専用モードと書き込み専用モードがある
	アクセス権限を制御したり誤って書き込むことを防止するためにモードに分ける

	書き込みモードはos.Create("ファイルパス")で得れる
	新規作成もできるし既存のファイルにも使える
*/