package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)
    
    cnt := 0
    for i := 0 ; i < n ; i++ {
        var p, q int
        fmt.Scan(&p, &q)
        if (q - p) >= 2 {
            cnt++
        } 
    }
    fmt.Println(cnt)
}
