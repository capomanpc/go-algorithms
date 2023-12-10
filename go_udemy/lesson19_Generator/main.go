package main

import "fmt"

//関数
//ジェネレーターの実装

//ジェネレーターとは何らかのルールに従って連続した値を返し続ける仕組みのこと

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherints := integers()
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())

	//同じ関数でも別々の変数にそれぞれクロージャを参照することで、それぞれ独立して出力することが可能

}
