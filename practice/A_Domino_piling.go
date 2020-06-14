package main

import "fmt"

func find(m, n int) int {
    if m == 0 || n == 0 {
        return 0
    }
    if m == 1 {
        return n/2
    }
    if n == 1 {
        return m/2
    }
    max1 := m * (n/2) + find(m, n%2)
    max2 := n * (m/2) + find(n, m%2)
    
    if max2 > max1 {
        max1 = max2
    }
    return max1
}
func main() {
    var m int
    fmt.Scan(&m)
    var n int
    fmt.Scan(&n)

    fmt.Println(find(m, n))    
}
