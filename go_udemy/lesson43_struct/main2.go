package main

import (
	"fmt"
)

//匿名構造体について
/*
	匿名構造体とはインスタンス名が匿名なのではなく、型名が匿名である構造体のこと
	匿名構造体は型の定義と要素を同時に行う
	struct{
		a int
		b string
	}{
		a: 21,
		b: "abc",
	}

*/

func main() {
	post := struct {
		Title   string
		Content string
	}{
		Title:   "Go言語の匿名構造体の使い方",
		Content: "Go言語で匿名構造体を使用することで、簡単に一時的なデータ構造を作成できます。",
	}

	fmt.Println("タイトル:", post.Title)
	fmt.Println("内容:", post.Content)
}
