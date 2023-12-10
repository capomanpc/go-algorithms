package main

import "fmt"

//interface
//最もポピュラーな使い方。異なる型に共通の性質を付与する。

type Stringfy interface {
	ToString() string
}

/*
Stringfy interfaceの宣言
ToString()メソッドを持つものは全てStringfy型にもしてあげる、という意味
stringは戻り値、つまりCarとPersonのToString関数のreturnがstring型であるという事

*/

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v,Age=%v", p.Name, p.Age)
}

//(p *Person)構造体の引数
//ToString()メソッドの名前、変数の引数なし
//string　返り値がstring

/*
Sprint関数について

文字列を出力する関数
string print　の略
文字列を変数などに出力している

例
text := fmt.Sprint("Hello")

fmt.Print("Hello")
例えばこの関数は標準出力として画面に出力される
標準出力とはデフォルトの出力経路のことであり、その経路の先につながっているのは一般的に画面である

*/

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v,Model=%v", c.Number, c.Model)
}

//Car構造体、Person構造体それぞれについて構造が同じなToString()関数をふたつ用意している
//何故ふたつ用意しているかと言うとCar型とPerson型は型が違うのでToString関数を一つにまとめられないからである
//これを改善するためにinterface型を用いる

func main() {
	vs := []Stringfy{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "123-456", Model: "AB-1234"},
	}

	//Stringfy型インターフェースのスライスに構造体PersonとCarを入れている
	//Stringfy interface はメソッドの集まりだがスライスの要素は構造体や構造体のアドレスが含まれる、これは気にせず暗記することにする

	for _, v := range vs {
		fmt.Println(v.ToString())
	}
	//vは構造体のアドレス
	//構造体.メソッド名で構造体を指定してメソッドを実行
}
