package main

import "fmt"

//ラベル付きfor
//複数のfor文を一気に抜けたいときなどに使う

func main() {

	/*
			   Loop:
			   	for {
			   		for {
			   			for {
			   				fmt.Println("START")
			   				break Loop
			   			}
			   			fmt.Println("処理をしない")
			   		}
			   		fmt.Println("処理をしない")
			   	}
			   	fmt.Println("END")

		//breakはfor文が入っている一番手前のforだけ抜ける
		//breakの位置で一気にfor文を抜けたい場合Loopというラベルを抜けたいfor文の上に記述することで抜けられる

	*/
Loop:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue Loop
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("処理をしない")
	}

	//Loopラベルを付けることでラベルのすぐ下にあるfor文の先頭まで戻ることができる
	//Loopラベルの下にあるfor文の残りの処理を飛ばす

	//continue文は今位置しているfor文の残りの処理を飛ばす
	//今位置しているfor文の先頭に戻る
}
