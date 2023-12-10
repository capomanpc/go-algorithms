package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//ioutil
//I/O utilityの略
//簡易的な読み書きに関するパッケージ
//ファイルの読み書き

func main() {
	//入力全体を読み込む

	f, _ := os.Open("foo.txt")  //os.Open()でファイルハンドラ取得 and ファイル開く
	bs1, _ := ioutil.ReadAll(f) //読み込んだデータをbyteのスライスに格納、データが大きすぎる場合はbyteのスライスなので適していない
	//io.Readerインターフェースを持つ実装する任意のオブジェクトに対して使用される
	//非常に汎用的であり沢山のデータソースに対して使われる

	fmt.Println(string(bs1)) //byteスライスを文字列に変換して表示

	if err := ioutil.WriteFile("bar.txt", bs1, 0666); err != nil {
		log.Fatalln(err)
	}
	//bar.txtにbs1を書き込む、新しくファイルを作成する場合のファイル権限は0666
	//ioutil.WriteFile()は書き込みに失敗したときだけerror型を返す
	//指定したファイルがなければ新しくファイルを作成する

	bs2, _ := ioutil.ReadFile("foo.txt")
	fmt.Println(string(bs2))
}
