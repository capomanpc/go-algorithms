package main

import (
	"log"
	"os"
)

//log

func main() {
	//ログの出力先を標準出力に変更
	log.SetOutput(os.Stdout)

	//アプリケーションの動作情報、診断情報、エラー情報などを記録する際に使用される
	log.Print("Log\n") //デフォルトで日時を表示する
	log.Println("Log2")
	log.Printf("Log%d\n", 3)

	/*
		//fatal 致命的な（致命的なエラーで続行不可能）
		log.Fatal("Log\n")　//os.Exit(1)が内部で実行される
		log.Fatalln("Log2")
		log.Fatalf("Log%d\n", 3)
	*/

	/*
		//panic状態にしてプログラムを強制終了させる
		log.Panic("Log\n")
		log.Panicln("Log2")
		log.Panicf("Log%d\n", 3)
	*/

	/*
		//このコメントアウトでは任意のファイルを作成し、出力先に指定している
		//os.Create ファイルの作成
		f, err := os.Create("test.log")
		if err != nil{
			return
		}

		//作成したio.Writer型のファイルを出力先に設定
		log.SetOutput(f)
		log.Println("ファイルに書き込む")

	*/

}
