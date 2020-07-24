package main
import (
    "fmt"
    "os/exec"
    "flag"
)


func main(){
    flag.Parse()
    fmt.Println(GetRSS("https://www.nhk.or.jp/r-news/podcast/nhkradionews.xml"))
}

func GetRSS(url string) (string, int){
    out, err := exec.Command("curl",url).Output()
    if err != nil {
        return "",-1
    }
    return string(out), 0
}
