package main

import (
	"flag"
	"fmt"
	// "strconv"

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

    window(*podRssData, url)
}

func window(paraData rssLib.RssData, url string) {
    // 表示のズレを防止
    runewidth.DefaultCondition.EastAsianWidth = false

    // tview.TextViewを返す
    newPrimitive := func(text string) tview.Primitive {
        return tview.NewTextView().
            SetTextAlign(tview.AlignCenter).
            SetText(text)
    }
    menu    := newPrimitive("Menu")
    main    := newPrimitive("Main content")

    grid := tview.NewGrid().
        SetBorders(true).
        AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, false)

    grid.AddItem(menu, 1, 0, 1, 1, 0, 100, false).
        AddItem(main, 1, 1, 1, 1, 0, 100, false)


    if err := tview.NewApplication().SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}

    // infoStr := "Infomation about the " + url + "\n" +
    //            "\n" +
    //            "Title     : " + paraData.Channel.Title + "\n" +
    //            "CopyRight : " + paraData.Channel.CopyRight + "\n" +
    //            "Category  : " + paraData.Channel.ItunesCategory.Text + "\n" +
    //            "\n" +
    //            strconv.Itoa(len(paraData.Channel.Item)) + "話のエピソードがあります"

    // app := tview.NewApplication()
    // textView := tview.NewTextView().
    //     SetRegions(true).
    //     SetChangedFunc(func() {
    //         app.Draw()
    //     })
    // go func() {
    //     fmt.Fprintf(textView, "%s", infoStr)
    // }()
    // textView.SetBorder(true)
    // if err := app.SetRoot(textView, true).SetFocus(textView).Run(); err != nil {
    //     panic(err)
    // }

}

