package main

import "fmt"

type T struct {
	User User //UserフィールドにUser型構造体を埋め込む
}

/*
上記の埋め込みを以下のように省略できる

type T struct{
	User
}

fmt.Println(t.Name)
//user1

埋め込みでUserを省略した際、User型構造体のNameフィールドにアクセスしたかったら
上記のようにt.Nameのように省略できる
*/

type User struct {
	Name string
	Age  int
}

func (u *User) SetName() {
	u.Name = "A"
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}
	//T型の構造体のUserフィールドにUser型の構造体を定義
	//tはT型構造体のインスタンス名

	fmt.Println(t)
	//{{user1,10}}

	fmt.Println(t.User)
	//{user1,10}

	fmt.Println(t.User.Name)
	//user1

	t.User.SetName()
	//埋め込んだ構造体に対してメソッドを使いたい場合上記にように記述することでメソッドを使用できる

	fmt.Println(t)

}
