package main

import (
	"log"

	"github.com/Kawaeugtkp/myniceprogram/helpers" // さっきダメやったのに
	// なんかpreferenceの設定とかをやらずともうまくいっている。
	// その間にやったことといえばgo mod tidyをterminalで実行してみたことくらいだが
	// これが機能しているとはあんまり思えない
)

func main() {
	log.Println("Hello")
	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	log.Println(myVar.TypeName)
}