package main

import (
	"fmt"
	"net/http"
)

//ハンドラ...HTTPリクエストに対して応じるための機能 以下はハンドラなどについて
// https://hackmd.io/@rhHzPg4WS26yfiXdOaOMTg/B1mbxaGB8

type MyHandler struct{}

//http.ListenAndServe()の第二引数ではhttp.Handlerインタフェースを実装している必要があるため、まずは構造体の宣言

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!") //Hello Worldとwに書き込むメソッド
}

//http.HandlerインターフェースはServeHTTPメソッドの実装を要求するためMyHandler構造体でServeHTTPメソッドを実装する
/*
	func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)について

	・ServeHTTP メソッドは HTTP レスポンスのボディを作成するメソッド
	・Fprintfの出力先(第一引数)はio.Writerインタフェースを実装する任意の型でなければいけない
	　例えばos.File型構造体はWriteメソッドを持つのでio.Writerインターフェースを実装しており第一引数に入れることができる
	・http.ResoponseWriterインターフェースはWriteメソッドを有しておりio.Writerインターフェースを満たす
	　Fprint関数の第一引数はio.Writerを満たす任意の型であるのでhttp.ResoponseWriterインターフェースに指定できる

	http.ResponseWriterについて
	・レスポンスメッセージを構築するメソッドを持っているインターフェース
	・Writeメソッドを有しているのでio.Writerインターフェースを満たす
	・引数のwは構造体のインスタンスである

	http.Requestについて
	・リクエストメッセージが格納されている構造体


*/

func main() {

	//http.ListenAndServe(":8080", &MyHandler{})
	http.ListenAndServe(":8080", &MyHandler{})
	//サーバーを起動しリクエストを受信できる状態にしてリクエストに応じたハンドラを呼び出す関数

	/*
		http.ListenAndServeについて

		1.HTTPサーバーを開始して指定された第一引数のアドレス（この場合はポート8080）でリッスン状態にする
		リッスン状態...サーバーがリクエストを受け取れるような状態

		2.:8080の":"はポート番号を指定する標準的な方法
		"192.168.1.1:8080"のようにIPアドレスの後に付ける
		":"から始まる場合は「すべての利用可能なネットワークインターフェースのポート8080」と解釈される
		サーバーがそのマシンのすべてのIPアドレスでリクエストを待ち受けることを意味する

		3.第二引数が nil の場合、Go言語のHTTPサーバーはデフォルトのハンドラである http.DefaultServeMux を使用
		http.DefaultServeMuxはハンドラを適切に振り分けるハンドラ
		ハンドラとはHTTPリクエストに対して要求を提供する関数などのこと
		たとえhtmlファイルを公開する場合、ファイルシステムから規定のファイルを探すハンドラが使われる

		4.第二引数にはhttp.Handlerインターフェースを実装したハンドラが必要
		http.HandlerインターフェースはServerHTTPメソッドの実装を要求する
		interfaceの復習をすればアドレスを第二引数に入れるのが一般的なことがわかる

	*/

	/*
		http.ListenAndServe(":8080", handler)の動作について

		(1)指定されたアドレス(8080ポート)でHTTPサーバーを起動
		(2)クライアント(Webブラウザ)から来るHTTPリクエストの待機
		(3)HTTPリクエストを受信すると第二引数に渡されたハンドラ（この場合は handler）
		を使用してそのリクエストを処理する。handler は http.Handler インターフェイス
		を実装している必要がある。
		(4)リクエストに対してServerHTTPメソッドを自動的に呼び出す
		http.ListenAndServe関数がServeHTTPメソッドにhttp.ResponseWriter型の変数wと
		*http.Request型の変数rを引数として渡している
	*/

	/*
				Writeメソッドをインターフェースでまとめるメリット

				ファイルに対してのWrite()やバッファに対してのWrite()などはWriteメソッドであるため
				io.Writerインターフェースを満たす
				これらのWriteメソッドは各構造体で各々定義されているものである
				これらをio.Writerインターフェースでまとめる意味について解説する

				以下のコードはio.Writerインターフェースを引数に取って様々な出力先に対応できる関数を作っている
				io.Writerインターフェースがないと各々の出力先に各々の関数が必要になるが
				インターフェースのおかげで一つの関数にまとめることが出来ている


				func WriteHelloWorld(w io.Writer) error {
		    		_, err := w.Write([]byte("Hello, World!"))
		    		return err
				}

				func main() {
		    		// ファイルへの書き込み
		    		file, _ := os.Create("hello.txt")
		    		defer file.Close()
		    		WriteHelloWorld(file)

		    		// バッファへの書き込み
		    		buf := new(bytes.Buffer)
		    		WriteHelloWorld(buf)

		    		// 標準出力への書き込み
		    		WriteHelloWorld(os.Stdout)
				}



	*/
}
