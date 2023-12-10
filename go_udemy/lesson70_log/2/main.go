package main

import (
	"log"
	"os"
)

//log

func main() {
	log.SetOutput(os.Stdout)
	//output先をosのstandard outputにsetしている

	//logのフォーマット(形式)を指定するプログラム達
	//標準のログフォーマット

	log.SetFlags(log.LstdFlags)
	//logのflag(オプションのこと)をスタンダードフラグ(標準のオプション)にセットしている
	//logに日時を追加する設定をしている
	//LstdFlags...Log Standard Flags
	log.Println("A")

	//マイクロ秒を追加
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	//日付且つ時間且つマイクロ秒のフラグ(オプション)をセット
	log.Println("B")

	//時刻とファイルの行番号(短縮系)
	//Lshortfile...ファイル名と行番号を表示、パス名を表示するのでは無く短縮してファイル名だけ表示
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("C")

	log.SetFlags(log.LstdFlags)
	//フラグを標準に戻している

	//ログのプリフィックスを設定
	//プレフィックス prefix 接頭辞
	//表示されるlogの先頭に[LOG]を追加できる

	log.SetPrefix("[LOG]")
	log.Println("E")
}
