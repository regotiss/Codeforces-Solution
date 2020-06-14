package main

import "fmt"

func main() {
    var t int
    fmt.Scan(&t)
    
    tot := 0
    for i := 0 ; i < t; i++ {
        cnt := 0
        for j := 0; j < 3; j++ {
            var n int
            fmt.Scan(&n)
            if n == 1 {
                cnt++
            }
        }
        if cnt >= 2 {
            tot++
        }
    }
    fmt.Print(tot)
}
