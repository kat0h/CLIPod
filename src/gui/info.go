package gui

import (
	"./rssLib"
	"github.com/rivo/tview"
	"strconv"
)

type Info struct {
	*tview.TextView
}

func NewInfo() *Info {
	view := tview.NewTextView().SetDynamicColors(true)
	view.SetBorder(true).SetTitle("Info").SetTitleAlign(tview.AlignLeft)
	appview := &Info{TextView: view}
	return appview
}

func (a *Info) UpdateView(dat *rssLib.RssData, url string) {
	info := "Infomation about the " + url + "\n" +
		"\n" +
		"Title     : " + dat.Channel.Title + "\n" +
		"CopyRight : " + dat.Channel.CopyRight + "\n" +
		"Category  : " + dat.Channel.ItunesCategory.Text + "\n" +
		"\n" +
		strconv.Itoa(len(dat.Channel.Item)) + "話のエピソードがあります"
	a.SetText(info)
}
