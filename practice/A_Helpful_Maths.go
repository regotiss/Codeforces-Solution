package main

import (
    "fmt"
    "strings"
    "strconv"
)

func main() {
    var s string
    fmt.Scan(&s)
    
    n := make([]int, 3)
    for _, ns := range strings.Split(s, "+") {
        v, _ := strconv.Atoi(ns)
        n[v-1]++
    }
    var res []string
    for i, v := range n {
        for j := 0; j < v; j++ {
            res = append(res, string((i+1)+'0'))
        }
    }
    fmt.Println(strings.Join(res, "+"))
}
