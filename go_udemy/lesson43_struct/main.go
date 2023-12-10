package main

import "fmt"

//構造体

type User struct {
	Name string
	Age  int
	// X,Y int まとめて定義できる
}

//User型の構造体の定義
//string型のNameフィールド
//int型のAgeフィールド

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User
	fmt.Println(user1)
	//Nameの初期値は空文字でAgeの初期値は0

	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)
	//User{}とすることで初期値でのuser2を定義できる
	user2.Name = "user2"
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)
	//任意の初期値を持たせて宣言

	user4 := User{"user4", 40}
	fmt.Println(user4)
	//フィールド名を省略して宣言

	/*
		user5 := User{30,"user5"}
		fmt.Println(user5)
		//フィールド名を省略して宣言する場合フィールドの順番を揃えないとエラーになる
	*/

	user6 := User{Name: "user6"}
	fmt.Println(user6)
	//Nameフィールドだけ宣言

	user7 := new(User)
	fmt.Println(user7)
	//new関数を使うと構造体のポインタ型になる
	//出力結果は以下のように&がついてポインタ型となる
	//&{ 0}

	user8 := &User{}
	fmt.Println(user8)
	//adress演算子をつけて宣言するとnewと同じようにポインタ型で定義できる
	//adress演算子をつけて宣言することが一般的

	UpdateUser(user1)
	UpdateUser2(user8)

	fmt.Println(user1)
	fmt.Println(user8)

}

/*

構造体とポインタ
	明示的な定義
		var user *User //userという構造体のポインタを宣言
		user = &User{}　//userというポインタにUser型の構造体を割り当てている

		//golangでは一般的にインスタンス名は付けずにポインタから直接操作する
		//フィールドの値は宣言されていない
		//後から初期化したい場合や初期化されたものが欲しいときに使う

	暗示的な定義
		userPtr := &User{"alice",20}

		//User型のポインタの宣言をフィールド値の指定
		//インスタンス名は宣言されていない


*/
