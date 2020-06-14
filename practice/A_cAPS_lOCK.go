package main

import (
    "fmt"
    "strings"
)

func isCaps(c byte) bool {
    return c >= 'A' && c <= 'Z'    
}

func main() {
    var s string
    fmt.Scan(&s)
    
    rule1, rule2 := false, true
    if !isCaps(s[0]) {
        rule1 = true
    }  

    for i, _ := range s {
        if i == 0 {
            continue
        }
        if !isCaps(s[i]) {
            rule2 = false
        } 
    }
    if rule1 && rule2 {
        fmt.Println(strings.Title(strings.ToLower(s)))   
        return
    } 
    if rule2 {
        fmt.Println(strings.ToLower(s))   
        return
    } 
    fmt.Println(s)   
}
