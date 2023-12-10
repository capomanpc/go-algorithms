package main

import (
	"log"
	"os"
)

//log

func main() {
	//ロガーの生成
	//ロガーとはログ情報を生成、管理、および出力するための仕組みやパッケージ
	//logのフラグを使い分けたいときにlogパッケージでは不便
	//log.New()メソッドを利用して新しいloggerパッケージを作成する
	//そうすることでlogのフラグを使い分けることができる

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	//log.New(logの出力先,接頭辞,flags(ログのオプションのこと))
	logger.Println("message")
	log.Println("messge")

	/*
		//下記コードはlogが実際に使用されている例
		//エラーハンドリングで使用される

		//条件分岐　エラーで終了させる

		_, err := os.Open("fdafsafa") //存在しないファイルを開かせる
		if err != nil{

			//ログ出力
			log.Fatalln("Exit",err)
			//log出力して強制終了

			//logger.Fatalln("Exit",err)
		}

	*/
}
