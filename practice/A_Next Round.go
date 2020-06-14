package main

import "fmt"

func main() {
    var k, nl int
    fmt.Scan(&nl)
    fmt.Scan(&k)
    
    var n []int
    
    for i := 0; i < nl; i++ {
        var s int
        fmt.Scan(&s)
        n = append(n, s)
    }
    
    i := nl-1
    for ; i >= 0; i-- {
        if n[i] <= 0 {
            continue
        }
        if n[i] >= n[k-1] {
            break
        }
    }
    
    fmt.Println(i+1)
}
