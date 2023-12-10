package main

import (
	"bufio"
	"fmt"
	"os"
)

//buffered I/O
//バッファリングされた入出力をサポートする
//バッファリングは、データの読み書きを効率化し、パフォーマンスを向上させるために使用されるテクニック

func main() {
	//標準入力を行単位で読み込む

	//標準入力をソースにしたスキャナの生成
	scanner := bufio.NewScanner(os.Stdin)

	//入力のスキャンが成功する限り繰り返すループ
	for scanner.Scan() {
		//スキャン内容を文字列で出力
		fmt.Println(scanner.Text())
	}

	//スキャンにエラーが発生した場合の処理
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr,"読み込みエラー",err)
	}

}


/*
	解説
	scanner := bufio.NewScanner(os.Stdin)
	標準入力をソースにしてスキャナを生成している

	スキャナ
		入力（今回の場合は標準入力）から受け取ったデータを高度に
		処理したり解析したりするツールやオブジェクトのこと
		今回のコードでは行を読み取っている

	ソース
		データの入力元や情報の源
		os.Stdinにあたる
		標準入力やファイルからの入力などのこと

	scanner
		NewScanner()で作成したScanner型構造体のオブジェクト名
		スキャナのこと

	NewScanner
		Scanner型構造体からオブジェクトを生成する関数
	
	os.Stdin
		標準入力（ソース）のこと
	　
	　
	for scanner.Scan() {
		//スキャン内容を文字列で出力
		fmt.Println(scanner.Text())
	}

	scanner.Scan()がtrueかfalseを返すのでそれでforの判断をする
	生成したscannerオブジェクトで定義されているScan()で行ごとに読み取る
	scanner.Text()で読み取った文字列にアクセスしている



	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr,"読み込みエラー",err)
	}

	ループが終了した際にif文が実行される
	forがエラーで終了したときなどは、このif文が動作する
	err := scanner.Err();はif文のブロック内だけで有効な変数定義
	ifキーワードと条件式の間に変数定義を入れることができる
	scanner.Err()でエラーを取得している



*/
