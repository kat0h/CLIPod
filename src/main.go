package main

import (
    "flag"
    "fmt"

    // サブ
    "./rss"

    // 外部ライブラリ
    "github.com/mattn/go-runewidth"
    "github.com/rivo/tview"
)


func main(){
    // URLを読む
    flag.Parse()
    url := flag.Arg(0)
    // RSSの取得
    xmlData, retFlag := rss.GetRSS(url)
    if retFlag == -1{
        fmt.Println("エラー: URLを指定してください。")
        return
    }
    rssData := rss.XmlParse(xmlData)
    fmt.Printf("%#v\n", rssData)

    runewidth.DefaultCondition.EastAsianWidth = false
    window()
}

func window() {
    box := tview.NewBox().
        SetBorder(true).
        SetTitle("Box Demo")
        if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
            panic(err)
        }
}
