package main

import "fmt"

//型
//byte(unitB)型

//byte型は符号なしの8bit整数(0~256)を格納するデータ型

func main() {
	byteA := []byte{72, 73}
	//byte型の配列を作っている
	//i:=[]int{72,73}と同じようにbyte型は一つの文字を入れるデータ型

	fmt.Println(byteA)
	//角カッコ..スライスと呼ばれる型。配列としてデータを格納できる
	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Println(c)

	fmt.Println(string(c))

}
