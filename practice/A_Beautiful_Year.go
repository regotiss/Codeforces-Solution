package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)
    
    for {
       n++
       d1 := n / 1000
       d2 := n / 100 % 10
       d3 := n / 10 % 10
       d4 := n % 10
       
       if d1 != d2 &&  d1 != d3 && d1 != d4 && d2 != d3 && d2 != d4 && d3 != d4 {
           break
       }     
    }   
    fmt.Println(n)
}
