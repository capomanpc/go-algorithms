package main

import (
	"fmt"
	"time"
)

//time

func main() {
	//時間の生成
	//今の時間
	t := time.Now()
	fmt.Println(t)

	//指定した時間を生成

	//2020/06/10 10時10分10秒0ナノ秒
	//time.Localは実行環境のタイムゾーンを表す

	t2 := time.Date(2020, 6, 10, 10, 10, 10, 0, time.Local)
	fmt.Println(t2)
	fmt.Println(t2.Year())
	fmt.Println(t2.YearDay()) //通算日
	fmt.Println(t2.Month())
	fmt.Println(t2.Weekday()) //曜日
	fmt.Println(t2.Day())
	fmt.Println(t2.Minute())
	fmt.Println(t2.Second())
	fmt.Println(t2.Nanosecond())
	fmt.Println(t2.Zone())

}
