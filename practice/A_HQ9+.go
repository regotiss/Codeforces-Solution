package main

import "fmt"

func main() {
    var s string
    fmt.Scan(&s)
    
    for _, c := range s {
        if c == 'H' || c == 'Q' || c == '9' {
            fmt.Println("YES")
            return
        }
    }
    fmt.Println("NO")
}
