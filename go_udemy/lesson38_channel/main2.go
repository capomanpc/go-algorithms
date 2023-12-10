package main

import (
	"fmt"
	"time"
)

// この関数は、チャネルに1から10までの数字を送信する
func sendNumbers(channel chan int) {
	for i := 1; i <= 10; i++ {
		channel <- i // チャネルにiを送信する
	}
	close(channel) // チャネルを閉じる
}

// この関数は、チャネルから数字を受信して表示する
func receiveNumbers(channel chan int) {
	for n := range channel { // チャネルからデータが来るまで待つ
		fmt.Println(n) // チャネルから受信したnを表示する
	}
}

func main() {
	ch := make(chan int) // int型のチャネルを作成する

	go sendNumbers(ch)    // ゴルーチンでsendNumbers関数を実行する
	go receiveNumbers(ch) // ゴルーチンでreceiveNumbers関数を実行する

	time.Sleep(1 * time.Second)
	// main関数が終了してしまうとgoRoutineも終了してしまうので1秒待つ
}
