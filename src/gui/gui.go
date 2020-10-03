package gui

import (
	"./rssLib"
	"github.com/rivo/tview"
)

type Gui struct {
	Url      string
	Rss      *rssLib.RssData
	App      *tview.Application
	Info     *Info
	Episodes *Episodes
	Flex     *tview.Flex
}

func New(url string) *Gui {
	getRss := func(url string) *rssLib.RssData {
		xmlData, retFlag := rssLib.GetRSS(url)
		if retFlag == -1 {
			return new(rssLib.RssData)
		}
		retData := rssLib.XmlParse(xmlData)
		return retData
	}
	dat := getRss(url)
	g := &Gui{
		Url:  url,
		Rss:  dat,
		App:  tview.NewApplication().EnableMouse(true),
		Info: NewInfo(),
        Episodes: NewEpisodes(),
		Flex: tview.NewFlex().SetDirection(tview.FlexRow),
	}
	return g
}

func (g *Gui) Run() error {
	g.Info.UpdateView(g.Rss, g.Url)
    g.Episodes.UpdateView(g.Rss)
	g.Flex.AddItem(g.Info, 9, 1, false)
    g.Flex.AddItem(g.Episodes, 0, 1, true)

	if err := g.App.SetRoot(g.Flex, true).Run(); err != nil {
		return err
	}

	return nil
}
