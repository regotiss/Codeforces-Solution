package main

import (
    "fmt"
)

func main() {
    var n int 
    var s string
    fmt.Scan(&n, &s)
    
    cntA := 0
    for i := 0; i < n; i++ {
        if s[i] == 'A' {
            cntA++    
        } else {
            cntA--
        }
    }
    if cntA > 0 {
      fmt.Println("Anton") 
      return
    } 
    if cntA < 0 {
      fmt.Println("Danik") 
      return
    } 
    fmt.Println("Friendship")   
}
