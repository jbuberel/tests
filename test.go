package main

import (
    "fmt"
    "strings"
)

func substring(str, sub string) {
    fmt.Printf("%v\n", strings.Contains(str, sub))
}


func main(){
    var str = "Thompson"
    var sub = "son"
    substring(str, sub)
}
