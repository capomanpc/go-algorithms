package main

import (
	"fmt"
	"strconv"
)

//strconv
//基本的なデータ型とstring型の相互変換についてまとめられたパッケージ

func main() {
	//真偽値を文字列に変換する
	bt := true
	fmt.Printf("%T\n", strconv.FormatBool(bt))
	//%T...型を調べて表示する（%Type）
	//Format...データを特定の形式や構造に整形すること

	/*
		//整数を10進数の文字列に変換する
		//strconv.FormatInt(整数,基数)

		i := strconv.FormatInt(-100,10)
		fmt.Printf("%v,%T\n",i,i)
	*/

	/*
		//簡易的に変換できる
		//アイトゥエー関数 integer to ASCII 関数

		i2 := strconv.Itoa(100)
		fmt.Printf("%v,%T\n",i2,i2)
	*/

	/*
		//strconv.FormatFloat(実数値,フォーマット指定子,小数点以下の桁数,浮動小数点のビット数)
		//小数点以下の桁数を指定する引数で"-1"の場合、自動的に桁数が指定される
		//浮動小数点のビット数では32を指定するとfloat32の精度、64を指定するとfloat64になる

		//浮動小数点型文字列に変換する
		fmt.Println(strconv.FormatFloat(123.456,'E',-1,64))

		//指数表現による文字列化（小数点以下二桁まで）
		fmt.Println(strconv.FormatFloat(123.456,'e',2,64))

		//実数表現による文字列化
		fmt.Println(strconv.FormatFloat(123.456,'f',-1,64))

		//実数表現による文字列化（小数点以下二桁まで）
		fmt.Println(strconv.FormatFloat(123.456,'f',2,64))

		//指数部の大きさで変動する表現による文字列化（'g'は自動的に'f'か'e'かを判断する）
		fmt.Println(strconv.FormatFloat(123.456,'g',-1,64))
		fmt.Println(strconv.FormatFloat(123456789.123,'f',-1,64))

		//指数部の大きさで変動する表現による文字列化(仮数部全体が4桁まで)
		fmt.Println(strconv.FormatFloat(123.456,'g',4,64))

		//指数部の大きさで変動する表現による文字列化(仮数部全体が8桁まで)
		fmt.Println(strconv.FormatFloat(123456789.123,'G',8,64))

		'f': 123.456
		'e': 1.234560e+02
		'E': 1.234560E+02
		'g': 123.456 または 123.457
		'G': 123.456 または 123.457
		'b': 1111011.0111010011110001001001111100001010001111010111100


	*/

	/*
		//文字列を真偽値に変換する
		//trueへ変換できる文字列

		bt1, _ := strconv.ParseBool("true")
		fmt.Printf("%v,%T\n",bt1,bt1)

		bt2, _ := strconv.ParseBool("1")
		bt3, _ := strconv.ParseBool("t")
		bt4, _ := strconv.ParseBool("T")
		bt5, _ := strconv.ParseBool("TRUE")
		bt6, _ := strconv.ParseBool("True")
		fmt.Println(bt2,bt3,bt4,bt5,bt6)

		//これら全ての文字列をbool値(true)に変換できる


		//falseに変換できる文字列
		//二番目の戻り値はerror型なので、エラーハンドリングも可能
		bf1, ok := strconv.ParseBool("false")
		if ok != nil{
			fmt.Println("Convert error")
		}

		fmt.Printf("%v,%T\n",bf1,bf1)
		bf2, _ := strconv.ParseBool("0")
		bf3, _ := strconv.ParseBool("f")
		bf4, _ := strconv.ParseBool("F")
		bf5, _ := strconv.ParseBool("FALSE")
		bf6, _ := strconv.ParseBool("False")
		fmt.Println(bf2,bf3,bf4,bf5,bf6)


	*/

	/*
		//文字列を整数型に変換する
		//strconv.ParseInt(整数に変換する文字列, 基数, ビット数の精度)
		//ビット数の精度で0を指定した場合はgoのint型の精度（int64）が指定される

		i3, _ := strconv.ParseInt("12345", 10, 0)
		fmt.Printf("%v, %T\n", i3, i3)
		i4, _ := strconv.ParseInt("-1", 10, 0)
		fmt.Printf("%v, %T\n", i4, i4)

		//簡易的に変換できる

		i10, _ := strconv.Atoi("123")
		fmt.Println(i10)
	*/

	/*
		//文字列を浮動小数点数に変換する
		//strconv.ParseFloat(floatに変換る文字列,精度)

		fl1, _ := strconv.ParseFloat("3.14",64)
		fl2, _ := strconv.ParseFloat(".2",64)
		fl3, _ := strconv.ParseFloat("-2",64)
		fl4, _ := strconv.ParseFloat("1.2345e8",64)
		fl5, _ := strconv.ParseFloat("1.2345E8",64)
		fmt.Println(fl1,fl2,fl3,fl4,fl5)
	*/

}
