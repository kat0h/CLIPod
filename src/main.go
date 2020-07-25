package main
import (
    "fmt"
    "os/exec"
    "flag"
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
    out, err := exec.Command("curl",url).Output()
    if err != nil {
        return "",-1
    }
    return string(out), 0
}
