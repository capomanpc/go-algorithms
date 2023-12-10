package main

import (
	"fmt"
	"sync"
	"time"
)

//コードが分からなかったらフォルダ1を見る
/*
	ミューテックスによる排他処理
	ミューテックス Mutex -Mutual Exclusionの略,直訳すると相互排他

	複数のプロセスやスレッドが1つのリソースを利用するときに処理が被らないようにする機能のこと
	リソースは1つしかないので同時に利用するとrace condition(競合状態)になる
	ここで言うリソースはデータベース、メモリ、ファイルシステムなどを指す

	今回のコードの共有リソースはst2構造体のこと

	プロセス...main関数で動くプログラムのこと
	スレッド...main関数とは別に自立して動くことの出来る関数
	スレッドはプロセス内で動作するんですが、main関数の処理順とは関係なく動くことが出来ます
*/

var st2 struct{ A, B, C int }

var mutex *sync.Mutex //sync.Mutex 型のポインタ mutex を宣言
//mutex変数はmain関数の外でも使用するためグローバル変数として宣言する
//場合によってはmain関数で宣言して他の関数にポインタ型として送るのもok

func UpdateAndPrint(n int) {
	mutex.Lock() //mutexをロック
	//こうすることで他のスレッド(他のゴルーチン)がst2構造体に変数nを代入できなくなる
	//ロックすると処理を独占できるイメージ

	st2.A = n
	time.Sleep(time.Microsecond)
	st2.B = n
	time.Sleep(time.Microsecond)
	st2.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st2)

	mutex.Unlock() //mutexをアンロック
	//アンロックすることで他のスレッド(他のゴルーチン)がアクセスできなくなる
}

func main() {
	mutex = new(sync.Mutex) //mutex変数を初期化,new()関数と何故main関数で初期化するかは後で調べる
	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		}()
	}
	for {
	}
}
