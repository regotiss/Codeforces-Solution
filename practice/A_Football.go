package main

import "fmt"

func main() {
    var s string
    fmt.Scan(&s)
    
    if len(s) < 7 {
        fmt.Println("NO")
        return
    }
    cnt := 1
    var flg bool
    for i := 1; i < len(s); i++ {
        if s[i] != s[i-1] {
            cnt = 1
            continue
        } 
        cnt++
        if cnt >= 7 {
            flg = true
            break
        }
        
    }
    if flg {
        fmt.Println("YES")    
        return
    } 
    fmt.Println("NO")    
}
