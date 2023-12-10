package main

import "fmt"

//cahnnel
//複数のゴルーチン間でのデータの受け渡しをするために設計されたデータ構造
//宣言、操作

func main() {

	var ch1 chan int
	//双方向のチャネル
	//int型のデータの保持をできるchannel宣言
	//これはnillのchannelを宣言しているので書き込みや読み込みができない

	//var ch2 <-chan int
	//受信専用のチャネル

	//var ch3 chan<- int
	//送信専用のチャネル

	ch1 = make(chan int)
	//make関数を使う事でchannelの生成と初期化を行ってデータの書き込みや読み込みができる状態にできる

	ch2 := make(chan int)
	//宣言も

	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))
	//cap関数を用いることでバッファサイズを確認できる

	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))
	//容量を指定して宣言

	//チャネルはデータ送受信を行うデータ型なのでチャネルでデータを受信したりチャネルでデータを送信したりできる

	ch3 <- 1
	//ch3に1というint型のデータを送信

	fmt.Println("len", len(ch3))
	//データ数を表示するlen関数

	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3))

	//channelからデータを受信する場合
	i := <-ch3
	fmt.Println(i)
	//1
	fmt.Println("len", len(ch3))

	i2 := <-ch3
	fmt.Println(i2)
	//2
	fmt.Println("len", len(ch3))

	fmt.Println(<-ch3)
	//3
	fmt.Println("len", len(ch3))

	//channelのデータ構造はキュー、先に入れたものが先に出てくる、積み重ねて下から出てくる

	//バッファサイズ（5）を超えた場合はどうなるか

	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6

	//バッファサイズを超えた場合はデッドロックが掛かってエラーになる
}
