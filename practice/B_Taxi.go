package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    
    cnt := make(map[int]int)
    for i := 0 ; i < n ; i++ {
        var s int
        fmt.Scan(&s)
        cnt[s]++
    }
    ans := cnt[4] + cnt[3] + cnt[2]/2 
    cnt[1] -= cnt[3]
    
    if cnt[2] % 2 > 0 {
        cnt[1] -= 2
        ans++
    }
    
    if cnt[1] > 0 {
        ans  += cnt[1] / 4
    }
    
    if cnt[1] % 4 > 0 {
        ans++
    }

    fmt.Println(ans)
}
