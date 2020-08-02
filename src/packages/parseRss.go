package parseRss

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

// RSSフィードの取得
// 引数
//    url : string rssのurlをしてい
// 戻り値
//    string : 結果
//    int    : エラ〜コード
func GetRSS(url string) (string, int){
    out, err := http.Get(url)
    if err != nil {
        return "",-1
    }
    body, _ := ioutil.ReadAll(out.Body)
    defer out.Body.Close()
    return string(body), 0
}

type rss struct {
    Channel Channel `xml:"channel"`
}
type Channel struct {
    Title           string           `xml:"title"`
    Link            string           `xml:"link"`
    Language        string           `xml:"language"`
    CopyRight       string           `xml:"copyright"`
    ItunesAuthor    string           `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd author"`
    Description     string           `xml:"description"`
    ItunesCategory  ItunesCategory   `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd category"`
    ItunesImage     ItunesImage      `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd image"`
    LastBuildDate   string           `xml:"lastBuildDate"`
    ItunesKeywords  string           `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd keywords"`
    Item            []Item           `xml:"item"`
}
    type ItunesCategory struct {
        Text            string   `xml:"text,attr"`
    }
    type ItunesImage    struct {
        Href            string   `xml:"href,attr"`
    }
type Item    struct{
    Title               string    `xml:"title"`
    Enclosure           Enclosure `xml:"enclosure"`
    PubDate             string    `xml:"pubDate"`
    Duration            string    `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd duration"`
}
    type Enclosure      struct {
        Url             string    `xml:"url,attr"`
        Length          string    `xml:"length,attr"`
        Type            string    `xml:"type,attr"`
    }
func XmlParse(xmlStr string) (*rss) {
    retData := new(rss)
    parseErr := xml.Unmarshal([]byte(xmlStr), &retData)
    if parseErr != nil {
        fmt.Println("XML Unmarshal error:", parseErr)
        return retData
    } else {
        return retData
    }
}