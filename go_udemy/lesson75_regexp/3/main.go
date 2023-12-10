package main

import (
	"fmt"
	"regexp"
)

func main() {
	//正規表現による文字列の分割
	//Sprit()正規表現パターンで分割する関数
	re1 := regexp.MustCompile((`(abc|ABC)(xyz|XYZ)`))
	fmt.Println(re1.Split("ASHVJV<HABCXYZKNJBJVKABCXYZ", -1)) //[ASHVJV<H KNJBJVK]
	//ABCXYZは削除して分割される

	re1 = regexp.MustCompile(`\s+`)
	fmt.Println(re1.Split("aaaaaaaaaa     bbbbbbb  cccccc", -1)) //[aaaaaaaaaa bbbbbbb cccccc]
}
