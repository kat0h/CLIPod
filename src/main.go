package main

import (
	"flag"
	"fmt"
    "strconv"

	// サブ
	"./rssLib"

	// Lib
	"github.com/mattn/go-runewidth"
	"github.com/rivo/tview"
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
    // テスト用(URLの指定がないとNHKを利用する)
    if url == "" {
        url = "https://www.nhk.or.jp/r-news/podcast/nhkradionews.xml"
    }
    // RSSの取得
    podRssData := getRssData(url)
    // 表示のズレを防止
    runewidth.DefaultCondition.EastAsianWidth = false

    infoStr := "Infomation about the " + url + "\n" +
               "\n" +
               "Title     : " + podRssData.Channel.Title + "\n" +
               "CopyRight : " + podRssData.Channel.CopyRight + "\n" +
               "Category  : " + podRssData.Channel.ItunesCategory.Text + "\n" +
               "\n" +
               strconv.Itoa(len(podRssData.Channel.Item)) + "話のエピソードがあります\n"
    print(infoStr)

    info := tview.NewTextView()
    fmt.Fprintf(info, "%s", infoStr)
    info.SetBorder(true)

    epis := tview.NewTextView()
    epis.SetBorder(true)

    app := tview.NewApplication()
    flex := tview.NewFlex().SetDirection(tview.FlexRow).
        AddItem(info , 0, 1, true).
        AddItem(epis, 0, 3, false)

    if err := app.SetRoot(flex, true).Run(); err != nil {
        panic(err)
    }
}


