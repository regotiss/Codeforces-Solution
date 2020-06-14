package main

import ( 
    "fmt"
    )

func main() {
    var n  int
    var s string
    fmt.Scan(&n, &s)
   
    r := string(s[0])
    cnt := 0
    for i, c := range s {
        if i == 0 {
            continue
        }
        if rune(r[len(r)-1]) == c {
            cnt++
            continue
        }
        r = r + string(c)
    }
    fmt.Print(cnt)
}
