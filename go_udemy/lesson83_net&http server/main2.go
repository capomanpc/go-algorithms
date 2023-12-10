package main

import (
	"html/template"
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "Hello World111!")
}

func main() {
	http.HandleFunc("/top", top)
	/*
		http://localhost:8080/top のようにアクセスすることで第二引数に指定したハンドラ関数を使える
		ハンドラ関数はhttp.Handlerインターフェースを満たさなくてもよい
		複雑なカスタマイズをしたい場合ServeHTTPメソッドを使用し、
		そこまでの設定を必要としない場合はHandleFuncを使う

	*/
}
