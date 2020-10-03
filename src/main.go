package main

import (
	"flag"

	// サブ
	"./gui"

	// Lib
	"github.com/mattn/go-runewidth"
)

func main() {
	// URLを読む
	flag.Parse()
	url := flag.Arg(0)
	// テスト用(URLの指定がないとNHKを利用する)
	if url == "" {
		url = "https://www.nhk.or.jp/r-news/podcast/nhkradionews.xml"
	}
	// 表示のズレを防止
	runewidth.DefaultCondition.EastAsianWidth = false

	gui.New(url).Run()
}
