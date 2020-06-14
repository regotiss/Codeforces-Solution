package main

import "fmt"

func main() {
    var s, t string
    fmt.Scan(&s, &t)
    
    if len(s) != len(t) {
        fmt.Print("NO")
        return
    }
    size := len(t)
    for i := 0; i < size; i++ {
        if s[i] != t[size - i - 1] {
            fmt.Print("NO")
            return
        }
    }
    fmt.Print("YES")
}
