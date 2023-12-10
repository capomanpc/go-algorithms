package main

import "fmt"

//型
//整数型

func main() {
	var i int = 100

	/*
		 	int8	8bitの符号付き整数型
			int16	16bitの符号付き整数型
			int32	32bitの符号付き整数型
			int64	64bitの符号付き整数型

			環境依存で上記のどれかが採用される

	*/

	var i2 int64 = 200

	fmt.Println(i2 + 50)
	fmt.Println(i + 50)

	//fmt.Println(i+i2)

	//int型とint64型同士の演算はできない
	//例えint型が64bitだとしてもint型とint64型は別型扱いされる

	fmt.Printf("%T\n", i2)

	//%Tは書式指定子で型を表示してくれる

	fmt.Printf("%T\n", int32(i2))

	fmt.Println(i + int(i2))

}
