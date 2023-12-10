package main

import (
	"fmt"
	"time"
)

//channel
//チャネルとごルーチン

func reciever(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	//fmt.Println(<-ch1)

	go reciever(ch1)
	go reciever(ch2)
	//reciever関数はch1とch2に値が入ってくるまで待つ
	//channelは待つという仕様を持っている

	i := 0
	for i < 100 { // このfor文でch1とch2に値を送信する
		ch1 <- i
		ch2 <- i
		time.Sleep(50 * time.Millisecond)
		i++
	}

}

//for i < 100　でch1とch2に値を送信して、並行処理をしているreciever関数で受信、そして表示をする
//main関数ではchannelに値を送信したら50ミリ秒待つ
