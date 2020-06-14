package main

import (
    "fmt"
    "strings"
)

func main() {
    var s string
    fmt.Scan(&s)
    
    u := 0
    l := 0
    for _, c := range s {
        if c >= 'A' && c <= 'Z' {
            u++
            continue
        } 
        l++
    }
    if u > l {
        fmt.Println(strings.ToUpper(s))    
    } else {
        fmt.Println(strings.ToLower(s))    
    }
}
