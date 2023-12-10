package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
} //返り値は*User型
//User型を返すコンストラクタ関数
//構造体を創り出す関数、引数を与えるだけで構造体を作れる

func main() {
	user1 := NewUser("user1", 10)

	fmt.Println(*user1)
}
