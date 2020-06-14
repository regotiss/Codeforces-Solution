package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    
    for i:=0; i < n; i++ {
        var s string
        fmt.Scan(&s)
        if len(s) <= 10 {
            fmt.Println(s)
            continue
        }
        fmt.Println(fmt.Sprintf("%s%v%s", string(s[0]), len(s)-2, string(s[len(s)-1])))
    }
}
