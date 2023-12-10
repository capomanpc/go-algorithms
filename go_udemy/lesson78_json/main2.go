package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type A struct{}

type User struct {
	ID      int       `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       A         `json:"A"`
}

func (u *User) UnmarshalJSON(b []byte) error {
	type User2 struct {
		Name string
	}
	var u2 User2
	err := json.Unmarshal(b, &u2)
	if err != nil {
		fmt.Println(err)
	}
	u.Name = u2.Name + "!"
	return err
}

/*
	UnmarshalJSONメソッドについて
	json.Unmarshalerインタ―フェースはfunc (u *User) UnmarshalJSON(b []byte) errorメソッドを要求する
	UnmarshalJSONメソッドが定義されている場合json.Unmarshal関数が呼び出されたとき、内部的にメソッドを呼び出す
	JSON形式から構造体に変換するときにデフォルトからカスタムをしたい場合このメソッドを定義する

	メソッドの動作について
	仮の構造体としてUser2型構造体を定義し、JSONの変換結果はUser2型構造体に入れる
	User2型構造体のデータに"!"をつけ足して、User型構造体のポインタ変数uに代入する
	構造体のポインタをメソッド内で変更することで戻り値に構造体を指定しなくても構造体を変更できる
*/

func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr " + u.Name,
	})
	return v, err
}

/*
	MarshalJSONメソッドについて
	json.marshalerインタ―フェースはfunc (u User) MarshalJSON() ([]byte, error)メソッドを要求する
	marshalJSONメソッドが定義されている場合json.marshal関数が呼び出されたとき、内部的にメソッドを呼び出す
	構造体からJSON形式に変換するときにカスタムしたい場合はこのメソッドを定義する

	メソッドの動作について
	匿名構造体をjsonに変換してバイト型スライスvに代入している
	json.Marshal関数の引数は構造体のポインタ型なのでアドレス演算子を付けて匿名構造体を指定している
	匿名構造体についてはstructのmain2.goを参照

*/

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "example@example.com"
	u.Created = time.Now()

	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bs)
	fmt.Println(string(bs))
}
