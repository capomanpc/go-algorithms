package main

import (
	"fmt"
	"time"
)

//channel
//close

func reciever(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
			//chの中身が空になる且つchがcloseしたらbreak
			//okという変数が中身が空になる且つchがcloseしたらfalseとなるため
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "end")
}

func main() {
	ch1 := make(chan int, 2)

	/*
		close(ch1)

		//ch1 <- 1

		//fmt.Println(<-ch1)

		i, ok := <-ch1
		fmt.Println(i, ok)

		//channelからの受信は二つの変数を出力することができる
		//iがchannelからの値
		//okがchannelがオープンかクローズされているかを示すbool値
		//厳密にはchannelのバッファ内が空であること且つchannelがクローズされている場合にokはfalseとなる
	*/

	go reciever("1.goroutine", ch1)
	go reciever("2.goroutine", ch1)
	go reciever("3.goroutine", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
	time.Sleep(3 * time.Second)
}
