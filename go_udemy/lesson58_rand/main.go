package main

import (
	"fmt"
	"math/rand"
	"time"
)

//rand

func main() {
	//乱数生成器の初期シードの設定
	//乱数生成器は実際にはランダムな数を生成しない
	//任意のシード値(256)に基づいて予測不可能に思える一連の数を生成
	//同じシード値を使うことで同じ乱数を再現することができる
	//シード値:初期値、開始値のこと。植物の種(seed)が由来

	rand.Seed(256)

	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())

	//現在の時刻をシードに使った疑似乱数の作成
	//シード値を現在の時刻の累積ナノ秒に設定している
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())

	//0-99の間の疑似乱数
	//イントエヌ関数　n未満のランダムな整数を生成する
	fmt.Println(rand.Intn(100))

	//int型の疑似乱数
	fmt.Println(rand.Int())

	//int32型の疑似乱数
	fmt.Println(rand.Int31())

	//int64型の疑似乱数
	fmt.Println(rand.Int63())

	//unit32型の疑似乱数
	fmt.Println(rand.Uint32())

}
