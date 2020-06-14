package main

import ( 
    "fmt"
    )


func main() {
    var n int
    fmt.Scan(&n)
    div  :=  []int{4,7,44,47,77,444,447,474,744,747,774,477,777}
    
    for _, d := range div  {
        if n % d == 0 {
            fmt.Println("YES")
            return    
        }      
    } 
    
    fmt.Println("NO")
}
