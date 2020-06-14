package main

import ( 
    "fmt"
    )


func main() {
    var s string
    fmt.Scan(&s)
    
    h := "hello"
    i := 0
    for _, c := range s {
        if i == len(h) {
            break
        }
        if c == rune(h[i]) {
            i++
        }
    } 
    if i == len(h) {
        fmt.Println("YES")   
        return
    }
    fmt.Println("NO")   
}
