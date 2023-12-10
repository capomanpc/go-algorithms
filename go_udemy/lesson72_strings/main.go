package main

import (
	"fmt"
	"strings"
)

//strings
//文字列操作をまとめたパッケージ

func main() {
	//文字列を結合する
	s1 := strings.Join([]string{"A", "B", "C"}, ",")
	s2 := strings.Join([]string{"A", "B", "C"}, "")
	//第1引数に文字列のスライス、第2引数でセパレータ(間に入れる文字)を指定
	fmt.Println(s1, s2)
	//A,B,C ABC　と出力される
	//スライスの中身をセパレータを使って左のように結合して表示してくれる

	//文字列に含まれる部分文字列を検索する
	i1 := strings.Index("ABCDE", "AB")
	i2 := strings.Index("ABCDE", "DE")
	//strings.Index(検索対象の文字列, 何文字目にあるか検索したい文字)
	fmt.Println(i1, i2) //0 3
	//検索対象の文字列に第二引数で指定した文字列が何文字目で出てくるかを出力する関数
	//部分文字列が検索対象に含まれていない場合-1を返す

	//検索対象の文字列に部分文字列"BC"が複数含まれているとき、最後に部分文字列が開始される位置を返す
	i3 := strings.LastIndex("ABCDABCD", "BC")
	fmt.Println(i3) //5

	//検索対象の文字列の中にA,B,Cのいずれかが始まる位置を出力
	//IndexAnyは最初の位置を返し、LastIndexは最後の位置を返す
	i4 := strings.IndexAny("ABC", "ABC")
	i5 := strings.LastIndexAny("ABC", "ABC")
	fmt.Println(i4, i5) //0 2

	//指定した文字列から開始されるかをbool値で返す
	//prefix 接頭辞　指定した接頭辞を持つ(Has)か判定
	b1 := strings.HasPrefix("ABC", "A")
	//指定した文字列で終わるかをbool値で返す
	//suffix 接尾辞
	b2 := strings.HasSuffix("ABC", "C")
	fmt.Println(b1, b2) //true true

	//指定した文字列が含まれているかをbool値で返す
	b3 := strings.Contains("ABC", "BC")
	//指定した文字列のいずれかが含まれているかをbool値で返す
	b4 := strings.ContainsAny("ABCDE", "BD")
	fmt.Println(b3, b4) //true true

	//第二引数の文字列が何回出現するか
	i6 := strings.Count("ABCABC", "B")
	i7 := strings.Count("ABCABC", "")
	fmt.Println(i6, i7) //2 7
	//空文字を渡した場合は全体の文字数をカウントする

	//文字列を繰り返して結合する
	//0を指定した場合は0回繰り返すので何も出力されない
	s3 := strings.Repeat("ABC", 4)
	s4 := strings.Repeat("ABC", 0)
	fmt.Println(s3, s4) //ABCABCABCABC
	//負の値を指定するとランタイムエラー

	//文字列の置換
	//指定した文字列を指定回数置換する
	//Replace(対象の文字列, 置換する対象の文字, 置換後の文字, 置換する回数)
	s5 := strings.Replace("AAAAAA", "A", "B", 1)
	s6 := strings.Replace("AAAAAA", "A", "B", -1) //-1と指定すると全ての文字を置換する
	fmt.Println(s5, s6)                           //BAAAAA BBBBBB

	//文字列を分割する
	//第二引数の文字で分割をして文字列型のスライスにして返す(セパレータを削除)
	s7 := strings.Split("A,B,C,D,E", ",")
	fmt.Println(s7)                                 //[A B C D E]
	s8 := strings.SplitAfter("A,B,C,D,E", ",")      //セパレータを残す
	fmt.Println(s8)                                 //[A, B, C, D, E]
	s9 := strings.SplitN("A,B,C,D,E", ",", 2)       //N分割するかを指定できる
	fmt.Println(s9)                                 //[A B,C,D,E]
	s10 := strings.SplitAfterN("A,B,C,D,E", ",", 4) //セパレータを残す　且つ　分割数を指定
	fmt.Println(s10)                                //[A, B, C, D,E]

	//大文字、小文字の変換
	s11 := strings.ToLower("ABC") //abc
	s12 := strings.ToLower("E")   //e

	s13 := strings.ToUpper("abc") //ABC
	s14 := strings.ToUpper("e")   //E
	fmt.Println(s11, s12, s13, s14)

	//文字列の前後から空白文字を取り除く　(文字列の前後には不要な空白が含まれるため)
	//空白文字　半角、全角、改行、復帰文字、タブなどもすべて削除
	s15 := strings.TrimSpace("    -    ABC    -    ") //文字列の前後のスペースを取り除く
	s16 := strings.TrimSpace("\tABC\r\n") //tab文字,復帰文字,改行も削除
	s17 := strings.TrimSpace("　　　　ABC　　　　") //文字列の前後の全角スペースも取り除く
	fmt.Println(s15, s16, s17)//-    ABC    - ABC ABC


	//文字列からスペースで区切られたフィールドを取り出してstringのスライスにして返す
	s18 := strings.Fields("a b c")
	fmt.Println(s18) //[a b c]
	fmt.Println(s18[1]) //b

}
