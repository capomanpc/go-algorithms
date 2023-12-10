package main

import (
	"fmt"
	"math"
)

//math

func main() {
	//小数点以下の切り捨て、切り上げ

	//math.Trunc 数値の正負を問わず、小数点以下は単純に切り捨てる
	//truncate 切り捨てる　トランケイト
	fmt.Println(math.Trunc(3.14))
	fmt.Println(math.Trunc(-3.14))

	//math.Floor 引数の数値より小さい最大の整数を返す
	//floor　床　数学的には下に丸めること
	//floor フロアと読む
	fmt.Println(math.Floor(3.5))
	fmt.Println(math.Floor(-3.5))

	//math.Ceir 引数の数値より大きい最小の整数を返す
	//ceiling 天井　数学的には上に丸めることｓ
	//Ceil シールと読む
	fmt.Println(math.Ceil(3.5))
	fmt.Println(math.Ceil(-3.5))

}
