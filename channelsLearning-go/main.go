package main

import (
	"log"

	"github.com/Kawaeugtkp/myprogram/helpers"
) // とりあえず自分で補完なしで
// ここを入力しました。これ関係はコード補完頼らないほうがいいのかもしれない

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan) // 関数の実行をずらして使えるのがdeferらしい。
	// 今回はintChanを閉じるというものらしいがこのclose関数についてはよくわかっていない

	go CalculateValue(intChan) // どうやらmain関数内でこのCalculateValueを
	// 呼んでもこの完了を待たずしてmainの部分が終わればプログラムは終了するらしい。
	// それではCalculateValueが使えないってことになるのでgoと表記することで
	// CalculateValueがちゃんと終わるのを待っているということ、、らしい

	num := <- intChan
	log.Println(num)
}