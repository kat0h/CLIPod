package main
import (
	// "encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)


func main(){
    flag.Parse()
    url := flag.Arg(0)
    xml, retFlag := GetRSS(url)
    if retFlag != -1{
        fmt.Println(xml)
    } else {
        fmt.Println("ERR")
    }
}

func GetRSS(url string) (string, int){
    out, err := http.Get(url)
    if err != nil {
        return "",-1
    }
    body, _ := ioutil.ReadAll(out.Body)
    defer out.Body.Close()
    return string(body), 0
}

func XmlParse(xml string){
    return
}
