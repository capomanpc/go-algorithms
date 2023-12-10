package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

//json
//構造体からJSONテキストへの変換
//エンコーダーとデコーダーの機能を持つ
/*
	jsonとは
	下記のようにkeyとvalueで表されたデータ構造のこと。配列などもある。
	{
  		"name": "Alice",
  		"age": 30,
  		"isStudent": false
	}
*/

type A struct{}

type User struct {
	ID      int       `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       A         `json:"A"`
	//この構造体がjsonに変換されたときフィールド名はkeyとなって表示されるが`json:"id"`などと追加することでidに変更できる
}

//AフィールドにA型構造体を埋め込んでいる
/*
	`json:"id"`について
	・構造体フィールドのタグと呼ばれるものであり、追加情報を表す
	・json:のあとに文字列を指定することでJSONでのkeyの表示方法を変えることができる

	`json:"id,string"`とすることでint型の整数をstring型にしてJSONに変換できる

	`json:"-"`とすることでJSONに変換したときkeyを消すことができる

	・構造体の各値が初期値(int型は0,string型は空文字)の場合で、初期値(0や空文字)をJSONで表示させたくない場合
	`json:"id,omitempty"` このようにomitemptyとすることで初期値を消すことができる
	omit...省略する empty...空(から)
*/

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

func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr " + u.Name,
	})
	return v, err
}

func main() {
	u := new(User) //User型構造体のuを作成
	u.ID = 0
	u.Name = ""
	u.Email = "example@example.com"
	u.Created = time.Now()
	//各フィールドに値を設定

	//Marshal 構造体をJSONに変換
	bs, err := json.Marshal(u)
	//bs...構造体をJSONに変換した値をbyteのスライスに格納
	//Marshallingが由来,意味はオブジェクトやデータ構造を他のフォーマットに変換するプロセス
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bs)
	fmt.Println(string(bs)) //JSON形式が表示される

	fmt.Printf("%T\n", bs) //型表示 byte[]型

	var u2 User

	//Unmarshal　JSONを構造体に変換
	//jsonデータが入ったbyte型スライスbsをUser型構造体のu2インスタンスに格納
	//第二引数に構造体のアドレスを指定
	//構造体のフィールド名とJsonのkey名が一致している必要がある
	//一致していればJSONを構造体に変換してくれる
	if err := json.Unmarshal(bs, &u2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(u2)
}
