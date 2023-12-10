package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//net/http
//クライアント

func main() {

	//GETメソッド
	//httpパッケージのGet関数
	//指定したURLに対してGETメソッドを行いレスポンスをresに格納
	//具体的にそのURLのhtmlファイルなどを取得するGETメソッドを送る
	res, _ := http.Get("")
	//http.Response型の構造体を返す、サーバーからのメッセージを構造体に格納

	fmt.Println(res.StatusCode) //200
	//ステータスコード 200,404など

	fmt.Println(res.Proto) // HTTP/2.0
	//プロトコルhttps://example.com

	//ここよくわからないから後でやる
	fmt.Println(res.Header["Date"])         //[Wed, 15 Nov 2023 04:48:53 GMT]
	fmt.Println(res.Header["Content-Type"]) //[text/html; charset=UTF-8]

	fmt.Println(res.Request.Method) //GET
	fmt.Println(res.Request.URL)    //https://example.com

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body) //HTTPメッセージをbyte型スライスに代入
	fmt.Print(string(body))

	/*
		Response BodyをCloseしないと、ファイルディスクリプタが枯渇するリスクがあります。
		tcpコネクションはpersistConnection構造体の中にnet.Connインターフェイスで保持しています。
		BodyをCloseしないと同構造体のclosedフラグがセットされず、net.ConnがCloseされません。
		すなわち、TCPコネクションがクローズされないために、そのまま続けてHttpリクエストを発行して
		いくとファイルディスクリプタが枯渇することになります。

		よくわからないからclose()する理由について上記をコピペして残しておく
		引用: https://qiita.com/stk0724/items/dc400dccd29a4b3d6471
	*/

	/*
		//------------------------------------
		//POSTメソッド

		vs := url.Values{}

		vs.Add("id", "1")
		vs.Add("message", "メッセージ")
		fmt.Println(vs.Encode()) // => "id=1&message=%E3%83%A1%E3%83%83%E3%82%BB%E3%83@<dtp>{lb}%BC%E3%82%B8"


		res, err := http.PostForm("https://example.com/", vs)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Print(string(body))
	*/

}

/*
	返ってきたhttpメッセージを格納するhttp.Response構造体の定義

	type Response struct {
    Status     string // ステータスライン（HTTPステータスコードとテキスト）
    StatusCode int    // HTTPステータスコード（整数）
    Proto      string // 使用されたHTTPプロトコル（通常は "HTTP/1.1" または "HTTP/2"）
    ProtoMajor int    // 使用されたHTTPプロトコルのメジャーバージョン
    ProtoMinor int    // 使用されたHTTPプロトコルのマイナーバージョン
    Header     Header // HTTPヘッダーのマップ
    Body       io.ReadCloser // レスポンスボディのリーダー
    ContentLength int64 // コンテンツの長さ（バイト単位）
    // 他にも多くのフィールドが含まれています
}


*/
