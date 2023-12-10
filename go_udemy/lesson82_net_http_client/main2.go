package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//HTTPのPOSTメソッドを使用して特定のURLにフォームデータを送信し、そのレスポンスを受け取っている

func main() {
	//POSTメソッド

	vs := url.Values{}
	//url.Values型は以下のように宣言されておりstring型のkeyと[]string型(stringのスライス)のvalueを持つmap型
	//type Values map[string][]string
	//mapの初期化サンプル vs := map[int]string{} keyがint型でvalueがstring型のmapの初期化

	vs.Add("id", "1")
	vs.Add("message", "メッセージ")

	/*
		こんな感じのmapになっている
		{
			"id": ["1"],
		    "message": ["メッセージ"]
		}
	*/
	fmt.Println(vs.Encode()) // => "id=1&message=%E3%83%A1%E3%83%83%E3%82%BB%E3%83@<dtp>{lb}%BC%E3%82%B8"
	//mapであるvsをクエリ文字列形式にエンコードする(上の文字列がクエリ文字列形式)
	//id=1とmessage=...の二つのペアが&で繋がれており、keyとvalueのセットが複数あるときに使われる
	//URLに含む事ができない文字(スペースや特定の文字)を他の文字で表し、%に続く二桁の16進数で表される
	//クエリはラテン語のquaerereが語源で「求める、問う」という意味

	res, err := http.PostForm("https://example.com/", vs) //httpのPOSTメソッドを発行
	//変数vsはPOSTのリクエストのボディ
	// resは*http.Response 型のレスポンスメッセージが格納

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()              //必ず閉じる、閉じる理由に関してはまた後で調べる
	body, _ := ioutil.ReadAll(res.Body) //byte型スライスに格納
	fmt.Print(string(body))
}
