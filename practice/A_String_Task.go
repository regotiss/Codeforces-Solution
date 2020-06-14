package main

import "fmt"

func main() {
    var s string
    fmt.Scan(&s)
    
    for _, c := range s {
        if c >= 'A' && c <= 'Z' {
            c = c+32
        }
        if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'y' {
            continue
        }
        fmt.Print(".", string(c))
    }
}
