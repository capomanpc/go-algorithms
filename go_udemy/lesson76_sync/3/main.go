package main

import (
	"fmt"
)

//ゴルーチンはmain関数が終わってしまうと同時に終了する
//ゴルーチンの起動自体は非常に高速であるため、一つ目のゴルーチンの出力もmain関数の終了に間に合わない

func main() {
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Goroutine")
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Goroutine")
		}
	}()
	/*
		for {

		}
	*/
}
