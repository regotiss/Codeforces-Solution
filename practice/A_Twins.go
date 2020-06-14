package main

import ( 
    "fmt"
    "sort"
)

func main() {
    var n int
    fmt.Scan(&n)
    
    var cs []int
    var tot  = 0
    for i := 0 ; i  < n ; i++ {
        var c int
        fmt.Scan(&c)
        
        cs = append(cs, c)
        tot += c
    }
    tot = tot/2
    
    sort.Ints(cs)
    
    ctot := 0
    i := n - 1
    for  ; i >=  0; i-- {
        ctot += cs[i]
        if ctot > tot  {
            break
        }
    } 
    fmt.Println(n-i)
}
