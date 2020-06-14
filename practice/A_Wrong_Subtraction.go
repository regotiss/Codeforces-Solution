package main

import "fmt"

func main() {
    var n, k int
    fmt.Scan(&n, &k)
    
    for k > 0 {
        if n % 10 == 0 {
            n = n / 10
        } else {
            n = n - 1
        }
        k--
    }
    fmt.Println(n)
}
