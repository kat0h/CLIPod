package main
import (
	"flag"
	"fmt"

    "pakcage/getRss.go"
)


func main(){
    flag.Parse()
    url := flag.Arg(0)
    xmlData, retFlag := GetRSS(url)
    if retFlag != -1{
        rssData := XmlParse(xmlData)
        fmt.Printf("%#v\n", rssData)
    } else {
        fmt.Println("ERR")
    }
}
