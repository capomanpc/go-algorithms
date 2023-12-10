package main

import (
	"fmt"
	"regexp"
)

func main() {

	re1 := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)

	//正規表現にマッチした文字列の取得
	//FindString 正規表現にマッチした最初の部分を文字列で返す
	fs1 := re1.FindString("AAAAABCXYZZZZ")
	fmt.Println(fs1) //ABCXYZ

	//FindAllString 正規表現にマッチした複数の文字列をまとめて文字列型のスライスで返す
	//第二引数で取得する文字列の最大数を指定できる
	//-1 とすれば取得した文字列全てを文字列型スライスに出力する
	fs2 := re1.FindAllString("ABCXYZABCXYZ", -1)
	fmt.Println(fs2) //[ABCXYZ ABCXYZ]

	//正規表現のグループによるサブマッチ
	//サブマッチとはマッチするプロセスそのものではなく、マッチしたグループ内の文字列のこと
	//FindAllStringSubmatch
	re1 = regexp.MustCompile(`(\d+)-(\d+)-(\d+)`) //`\d`は整数値、`+`は1回以上繰り返す
	//グループ化することでサブマッチを抽出したり再利用できる

	s := `
	0123-456-7889
	111-222-333
	556-787-899
	`

	ms := re1.FindAllStringSubmatch(s, -1)
	//"正規表現がマッチした文字列"と"サブマッチ"を以下のようなフォーマット返す
	//[[マッチした文字列　サブマッチ1 サブマッチ2 サブマッチ3]　[...省略]...省略]
	fmt.Println(ms) //[[0123-456-7889 0123 456 7889] [111-222-333 111 222 333] ...]
	//スライスの中にスライスがある

	for _, v := range ms {
		fmt.Println(v)
	} //スライスmsの要素はスライスなので、その要素となっているスライスをforで表示

	//以下のようにスライスmsの要素となっているスライスを表示できる
	fmt.Println(ms[0][0])
	fmt.Println(ms[0][1])
	fmt.Println(ms[0][2])

	fmt.Println(re1.ReplaceAllString("Tel: 000-111-222", "$3-$2-$1"))
	fmt.Println(re1.ReplaceAllString("Tel: 000-111-222", "あ-い-う"))
}
