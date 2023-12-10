package main

import "fmt"

type MyInt int

//int型をMyInt型として新しい型を作っている

func (mi MyInt) Print() {
	fmt.Println(mi)
} //メソッドも作ることができる

func main() {
	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)

	//i := 100
	//fmt.Println(mi + i)
	//型が揃ってないと計算ができない
}
