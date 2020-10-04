package gui

import(
    "./rssLib"
    "github.com/rivo/tview"
)

type Image struct {
    *tview.TextView
    image string
}

func NewImage() *Image {
    view := tview.NewTextView().SetDynamicColors(true)
    view.SetBorder(true)
    appview := &Image{TextView: view}
    return appview
}

func (a *Image) UpdateView(dat *rssLib.RssData) {
    imageurl := dat.Channel.ItunesImage.Href
    a.SetText(imageurl)
}
