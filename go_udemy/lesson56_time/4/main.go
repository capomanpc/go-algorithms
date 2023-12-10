package main

import (
	"fmt"
	"time"
)

//time

func main() {
	//指定時間のスリープ
	//5秒間停止
	time.Sleep(5 * time.Second) //goroutineなどで使用
	fmt.Println("5秒間停止後表示")
}
