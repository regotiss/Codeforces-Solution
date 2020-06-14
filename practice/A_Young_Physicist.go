package main

import ( 
    "fmt"
    )


func main() {
    var n int
    fmt.Scan(&n)
    
    sx, sy, sz := 0, 0, 0
    for i:=0 ; i < n ; i++ {
        var x, y, z int
        fmt.Scan(&x, &y, &z)
        sx += x
        sy += y
        sz += z
    }
    
    if sx == 0 && sy == 0 && sz == 0 {
        fmt.Println("YES")
        return
    }    
    fmt.Println("NO")
}
