package main

import (
	"fmt"
	"time"
)

func main() {
	//時刻の間隔を表現する
	//time.Duration型を返す
	//timeパッケージで定義されているtime.Duration型
	//時間の定数みたいなもん
	//duration　期間、持続時間

	fmt.Println(time.Hour)
	fmt.Printf("%T\n", time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)

	//time.Duration型を文字列から生成する
	d, _ := time.ParseDuration("2h30m")
	fmt.Println(d)

	t3 := time.Now()
	t3 = t3.Add(2*time.Minute + 15*time.Second)

	//time.Duration型はtime.Time型と一緒に使用することで威力を発揮する
	//time.Time型は実際の時刻を表す型
	//2*time.Minute + 15*time.Secondなどのtime.Duration型を足してあげることで
	//time.Time型を操作できる

	fmt.Println(t3)
	//現在時刻の2分15秒後を表す
}
