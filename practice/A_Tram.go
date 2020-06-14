package main

import ( 
    "fmt"
    )


func main() {
    var n int
    fmt.Scan(&n)
    
    cap := 0
    rem := 0
    for i := 0; i < n ; i++ {
       var a, b int
       fmt.Scan(&b, &a)
       rem = (rem - b) + a
       if cap < rem {
           cap = rem
       }
    }
    
    fmt.Println(cap)   
}
