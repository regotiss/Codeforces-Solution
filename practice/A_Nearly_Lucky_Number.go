package main

import "fmt"

func main() {
    var n int64
    fmt.Scan(&n)
    
    cnt := 0
    for  n > 0  {
        rem := n % 10
        if rem == 4 || rem == 7 {
            cnt++
        }
        n =  n / 10
    }
    if cnt == 0 {
        fmt.Println("NO")   
        return
    }
    
    for cnt > 0  {
        rem := cnt % 10
        if rem == 4 || rem == 7 {
            cnt = cnt / 10
            continue
        }
        break
    }
    if cnt > 0 {
        fmt.Println("NO")   
        return
    }
    fmt.Println("YES")   
}
