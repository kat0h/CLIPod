package gui

import (
	"./rssLib"
    "github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Episodes struct {
	*tview.Table
}

func NewEpisodes() *Episodes {
	view := tview.NewTable()
	view.SetTitle("Episodes").SetTitleAlign(tview.AlignLeft)
	view.SetBorder(true)
	eps := &Episodes{Table: view}
	return eps
}

func (a *Episodes) UpdateView(dat *rssLib.RssData) {
	a.SetCellSimple(0, 0, "Title")
	a.SetCellSimple(0, 1, "Duration")
	for i := 0; i < len(dat.Channel.Item); i++ {
		a.SetCellSimple(i+1, 0, dat.Channel.Item[i].Title)
		a.SetCellSimple(i+1, 1, dat.Channel.Item[i].Duration)
	}
    a.GetCell(1, 0).SetTextColor(tcell.ColorBlue)
    a.GetCell(1, 1).SetTextColor(tcell.ColorBlue)
}
