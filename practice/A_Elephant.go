package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    
    m := n / 5
    if n % 5 != 0 {
        m++
    }
    fmt.Println(m)
}
