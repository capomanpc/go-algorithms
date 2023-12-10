package main

import (
	"fmt"
	"sync"
)

//ゴルーチンの終了を待ち受ける
//今まではforの無限ループを付けることで実現していたがゴルーチンの終了を待ってくれる関数がある

func main() {
	wg := new(sync.WaitGroup) //sync.WaitGroup型のポインタを生成

	wg.Add(3) //待ち受けるゴルーチンの数を指定
	//待ち受けるゴルーチンの数は3

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
		wg.Done() //これで終わったかを確認
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Goroutine")
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Goroutine")
		}
		wg.Done()
	}()

	wg.Wait() //wg.Done()を三つ確認するまで待つ
}
