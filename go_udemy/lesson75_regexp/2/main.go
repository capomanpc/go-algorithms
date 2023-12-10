package main

import (
	"fmt"
	"regexp"
)

func main() {
	//正規表現による文字列の置換
	re1 := regexp.MustCompile(`\s+`)                      //一回以上繰り返されるスペース
	fmt.Println(re1.ReplaceAllString("AAA BBB CCC", ",")) //AAA,BBB,CCC
	//第一引数の文字列の中からre1にマッチする部分を第二引数に置換する

	re1 = regexp.MustCompile(`、|。`)                                 //、か。
	fmt.Println(re1.ReplaceAllString("私は、Golangを使用する、プログラマー。", "")) //私はGolangを使用するプログラマー

}
