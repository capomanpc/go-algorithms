package main

import "fmt"

//channel
//for

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)

	for i := range ch1 { //チャネルの要素の回数繰り返される
		fmt.Println(i)
	}

	/*
		rangeループはチャネルが受信可能である限り終了しない
		close()をすることでそのチャネルが受信できなくなりチャネルの中身が無くなったら自動的に終了する
		メインゴルーチンとは別に他のゴルーチンが開かれている場合、
		そのプログラムからチャネルへ送信される場合もあるので基本的にチャネルはclose()をすべき
	*/
}
