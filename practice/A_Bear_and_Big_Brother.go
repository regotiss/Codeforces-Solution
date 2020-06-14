package main

import ( 
    "fmt"
    )


func main() {
    var a, b int
    fmt.Scan(&a, &b)
    
    i := 1
    for {
        a = 3*a    
        b = 2*b
        if a > b {
            break
        }
        i++
    }
    
    fmt.Println(i)   
}
