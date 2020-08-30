package main

import (
	"flag"
	"fmt"
	"strconv"

	// サブ
	"./rssLib"

	// 外部ライブラリ
	"github.com/mattn/go-runewidth"
	// "github.com/rivo/tview"
)

func getRssData(url string) (*rssLib.RssData){
    xmlData, retFlag := rssLib.GetRSS(url)
    if retFlag == -1{
        fmt.Println("Error : url missing")
        return new(rssLib.RssData)
    }
    retData := rssLib.XmlParse(xmlData)
    return retData
}

func main(){
    // URLを読む
    flag.Parse()
    url := flag.Arg(0)
    if url == "" {
        url = "https://www.nhk.or.jp/r-news/podcast/nhkradionews.xml"
    }
    // RSSの取得
    podRssData := getRssData(url)

    window(*podRssData, url)
}

func window(paraData rssLib.RssData, url string) {
    runewidth.DefaultCondition.EastAsianWidth = false
    infoStr := "Infomation about the " + url + "\n" +
               "\n" +
               "Title     : " + paraData.Channel.Title + "\n" +
               "CopyRight : " + paraData.Channel.CopyRight + "\n" +
               "Category  : " + paraData.Channel.ItunesCategory.Text + "\n" +
               "\n" +
               strconv.Itoa(len(paraData.Channel.Item)) + "話のエピソードがあります"

    fmt.Printf("%s\n", infoStr)
}
