package main

import ( 
    "fmt"
    )

func main() {
    var s string
    fmt.Scan(&s)
   
    fnd := make(map[rune]bool, len(s))
    
    cnt := 0    
    for _, c := range s {
        if fnd[c] {
            continue
        }
        fnd[c] = true
        cnt++
    }

    if cnt%2 == 0 {
        fmt.Print("CHAT WITH HER!")   
        return
    }
    fmt.Print("IGNORE HIM!")
}
