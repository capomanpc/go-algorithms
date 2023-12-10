package main

import "fmt"

//定数-constant
//後から値を変更できない
//定数は他のパッケージでも使えるようにグローバルに書くことが多い(main関数の外に書く)

const Pi = 3.14
const (
	URL      = "http://xxx.co.jp"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

//値を省略することで前の定義をコピーできる

const (
	c0 = iota
	c1
	c2
)

//iotaは連続する数字を出力する

const Big = 942984984982989294 + 1

//グローバルに定義
//変数の頭文字を大文字にすると他のpackageからでも呼び出せるようになっている

func main() {
	fmt.Println(Pi)
	fmt.Println(URL, SiteName)
	fmt.Println(A, B, C, D, E, F)
	fmt.Println(Big - 1)
	fmt.Println(c0, c1, c2)

}
