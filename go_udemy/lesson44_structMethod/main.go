package main

import "fmt"

//struct
//メソッド
//任意の型に特化した関数を定義するための仕組み

type User struct {
	Name string
	Age  int
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

/*
	uはレシーバーの名前
	SayNameはメソッドの名前
	User型のNameを表示するメソッド

	レシーバとは受け取るインスタンスの名前のこと
	uはSayNameというメソッド内での受け取ったインスタンスの別名

	user1.SayName()

	これがメソッドの実行方法で、user1がレシーバ
*/

func (u User) SetName(name string) {
	u.Name = name
}

/*
メソッドと関数の比較
	(u User)は構造体の引数
	(name string)は変数の引数
	メソッドは関数と違って構造体専用の関数みたいなイメージ

*/

func (u *User) SetName2(name string) {
	u.Name = name
}

//一般的にポインタ型にしておくのがふさわしい
//引数となる構造体のuser1はポインタ型でなくてもok

func main() {
	user1 := User{Name: "user1"}
	user1.SayName()
	//メソッドを使う場合はレシーバー.メソッドとして使う

	user1.SetName("A")
	//SetNameメソッドは参照型ではないので構造体の中身をAに変更することはできない
	user1.SayName()

	user1.SetName2("A")
	user1.SayName()
	//一般的にポインタ型にしておくのがふさわしい

}
