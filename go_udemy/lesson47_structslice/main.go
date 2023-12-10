package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Users []*User

//*Userを要素とするUsersという名前のスライス型
//typeを使うことでスライスも型にできる

func main() {
	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}
	user4 := User{Name: "user4", Age: 40}
	//User型構造体のインスタンスをuser1,user2,user3,user4変数に保存

	users := Users{}
	//初期値でUsers型の定義,Users型は*Userを要素とするスライスの型
	//usersは*Userを要素とするUsers型(スライス型)の変数名

	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3, &user4)
	//users配列(Users型)にuser1からuser4を連結

	fmt.Println(users)

	for _, u := range users {
		fmt.Println(*u)
	}
	//スライスのrange for ループ
	//for index,value := range slice

	users2 := make([]*User, 0) //長さが0のスライスを宣言、後から要素を追加できる
	users2 = append(users2, &user1, &user2)

	for _, u := range users2 {
		fmt.Println(u)
	}

}
