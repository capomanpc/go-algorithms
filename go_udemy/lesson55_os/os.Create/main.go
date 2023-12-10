package main

import "os"

//os

func main() {
	//ファイル操作
	//Create
	f, _ := os.Create("foo.txt")
	//新規ファイルの作成、同じ名前のファイルがある場合は削除されて新しいのが作られる

	//fはファイルへのハンドルを表すos.File型の変数
	//ハンドルとはファイルへの参照のこと

	f.Write([]byte("Hello\n"))
	//バイトのスライスで書き込み

	f.WriteAt([]byte("Golang"), 6)
	//write at 指定された位置に書き込む
	//6byte目に書き込む

	f.Seek(0, os.SEEK_END)
	//ファイルの読み書きのカーソルを合わせるメソッド
	//f.Seek(基準位置から何バイト動かすか,基準位置)
	//基準位置のos.SEEK_ENDは末尾を表す

	f.WriteString("Yeah")
	//文字列をファイルに書き込むメソッド
}


/*
	os.Create()は書き込み権限が付いたが欲しい場合に使う
	Create()に既存のファイルパスを入力しても書き込み権限を得ることができる
	ただしファイルの中身は空になる

*/

/*
	書き込みをした後に defer f.Close() でファイルを閉じるのが普通だが
	書き込みが正常終了しなかった場合にはClose()がエラーを返す
	そのため書き込みをするときf.Write()を使う場合にはエラーハンドリングをするのがふさわしい

	無名関数を用いてdeferと共にf.Close()をするコードを以下に示す 

*/

/*
	f, err := os.Create("write.txt")
	if err != nil {
		fmt.Println("cannot open the file")
	}
	defer f.Close()
	defer func(){
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

// 以下write処理等を書く
*/
