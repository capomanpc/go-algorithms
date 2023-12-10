package main

import "fmt"

//スライス
//for

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	for i, v := range sl {
		fmt.Println(i, v)
	}

	/*
		i		index number
		v		value
		for i, v := range sl		rangeキーワードでスライスslのindex番号とvalueを取得し代入、繰り返し回数は要素があるところまで
	*/

	//index番号を使わない場合
	for _, v := range sl {
		fmt.Println(v)
	}

	//index番号だけ使う場合
	for i := range sl {
		fmt.Println(i)
	}

	//古典的for
	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}

}
