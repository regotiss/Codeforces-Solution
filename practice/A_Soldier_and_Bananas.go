package main

import ( 
    "fmt"
    )


func main() {
    var k,n,w int
    fmt.Scan(&k,&n,&w)
    
    am := k * (w*(w+1)/2)
    
    b := am - n
    
    if b <= 0 {
        b = 0
    }
    fmt.Println(b)   
}
