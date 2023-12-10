package main

import (
	"fmt"
	"time"
)

//time

func main() {
	//時刻の差分を取得

	t5 := time.Date(2021, 7, 24, 0, 0, 0, 0, time.Local)
	//日時をt5に指定して格納
	t6 := time.Now()
	fmt.Println(t6)

	//t5 - t6
	d2 := t5.Sub(t6)
	//t5との差分を表すsub関数
	fmt.Println(d2)

	//時刻を比較する
	//bool値が出力される
	fmt.Println(t6.Before(t5)) //t6はt5の前か？false
	fmt.Println(t6.After(t5))  //t6はt5の後か？True
	fmt.Println(t5.Before(t6))
	fmt.Println(t5.After(t6))

}
