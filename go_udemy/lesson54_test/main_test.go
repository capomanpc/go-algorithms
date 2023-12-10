package main

import "testing"

var Debug bool = false

func TestIsOne(t *testing.T) {
	//テストしたい関数はIsOneなのでこのような名前を付ける
	//Goでは名前がTestで始まる関数がテスト用の関数

	i := 0

	if Debug {
		t.Skip("スキップさせる")
	}
	//DebugがTrueならテストを行う
	//falseならテストをスキップする

	v := IsOne(i)
	//IsOneの出力がtrueであればテスト成功
	//IsOneの出力がfalseであれば下記if文に引っかかってエラー文が出力される

	if !v {
		t.Errorf("%v != %v", i, 1)
	}
}
