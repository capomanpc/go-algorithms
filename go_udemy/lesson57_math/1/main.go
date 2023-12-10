package main

import (
	"fmt"
	"math"
)

//math

func main() {
	//数学的な定数
	//円周率
	fmt.Println(math.Pi)

	//2の平方根
	fmt.Println(math.Sqrt2)

	//整数型に関する定数
	//float32で表現可能な最大値
	fmt.Println(math.MaxFloat32)

	//float32で表現可能な0ではない最小値
	fmt.Println(math.SmallestNonzeroFloat32)

	//float64で表現可能な最大値
	fmt.Println(math.MaxFloat64)

	//float64で表現可能な0ではない最小値
	fmt.Println(math.SmallestNonzeroFloat64)

	//int64 ver
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MinInt64)

}
