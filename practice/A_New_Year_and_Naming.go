package main

import "fmt"

func main() {
    var n,m int
    fmt.Scan(&n,&m)
    
    var ns []string
    for i := 0; i < n; i++ {
        var s string
        fmt.Scan(&s)
        ns = append(ns, s)
    }
    var ms []string
    for i := 0; i < m; i++ {
        var s string
        fmt.Scan(&s)
        ms = append(ms, s)
    }
    var q int
    fmt.Scan(&q)
    for i := 0; i < q; i++ {
        var y int
        fmt.Scan(&y)
        
        pn := (y-1) % n
        pm := (y-1) % m
        
        fmt.Println(ns[pn]+ms[pm])
    }
}
