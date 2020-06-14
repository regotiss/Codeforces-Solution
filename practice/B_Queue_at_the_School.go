package main

import "fmt"

func main() {
    var n, t int
    fmt.Scan(&n, &t)
    
    var s string
    fmt.Scan(&s)
    
    bs := []byte(s)
    for i := 0 ; i < t ; i++ {
        for j := 1; j < n; j++ {
            if bs[j] == 'G' && bs[j-1] == 'B' {
                bs[j] = 'B'
                bs[j-1] = 'G'
                j++
            }
        }
    }
    fmt.Println(string(bs))
}
