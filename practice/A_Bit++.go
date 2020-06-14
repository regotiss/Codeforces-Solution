package main

import (
    "fmt"
    "strings"
)

func main() {
    var t int
    fmt.Scan(&t)
    
    cnt := 0
    for i:=0; i < t; i++ {
        var s string
        fmt.Scan(&s)
        if strings.Contains(s, "+") {
            cnt++
            continue
        }
        cnt--
    }
    fmt.Println(cnt)
}
