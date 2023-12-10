package main

import (
	"fmt"
	"net/url"
)

//net/url

func main() {
	//URLを解析
	//parse 解析する、という意味
	//プログラミングの文脈では、この用語は通常、文字列やデータを取り、
	//それをより使いやすい形式や構造に変換するプロセス

	u, _ := url.Parse("http://example.com/search?a=1&b=2#top") //URL型構造体を生成する
	//uは構造体
	/*

			type URL struct {
		    Scheme   string		//httpやhttpsなどのプロトコル
		    Opaque   string    // エンコードされていないオパックなデータ
		    User     *Userinfo // ユーザ名とパスワードの情報
		    Host     string    // example.comなどの通常ドメイン名
		    Path     string		//サーバー上の特定のリソースへの具体的なパスを表します。
								//例: /search, /images/photo.jpg など。
		    RawPath  string    // エンコードされたpath情報
		    ForceQuery bool    // クエリが空でも "?" を付けるかどうか
		    RawQuery string    // エンコードされたクエリ文字列クエリとはwebサーバーに送る追加情報
		    Fragment string    // ページ内の特定のセクションや位置に直接ジャンプするためのもの
								//webページのトップなどの位置
		}
	*/
	fmt.Println(u.Scheme)   //http
	fmt.Println(u.Host)     //example.com
	fmt.Println(u.Path)     // /search
	fmt.Println(u.RawQuery) //a=1&&b=2
	fmt.Println(u.Fragment) //top

	fmt.Println(u.IsAbs()) //true
	fmt.Println(u.Query()) //map[a:[1] b:[2]]

	//URLを生成
	url := &url.URL{}
	url.Scheme = "https:"
	url.Host = "google.com"
	q := url.Query()
	q.Set("q", "Golang")
	url.RawQuery = q.Encode()

	fmt.Println(url)

}
