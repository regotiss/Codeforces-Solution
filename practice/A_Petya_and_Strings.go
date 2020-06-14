package main

import "fmt"

func lower(c byte) byte {
    if c >= 'A' && c <= 'Z' {
        c += 32
    } 
    return c; 
}
func main() {
    var s1, s2 string
    fmt.Scan(&s1, &s2)
    
    res := 0
    for i, _ := range s1 {
        c1 := lower(s1[i])
        c2 := lower(s2[i])
        if c1 > c2 {
            res = 1
            break
        } 
        if c1 < c2 {
            res = -1
            break
        }
    }
    fmt.Println(res)
}
