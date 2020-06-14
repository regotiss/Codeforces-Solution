package main

import (
    "os"
    "bufio"
    "fmt"
    "strconv"
)

var scanner *bufio.Scanner
func read() string {
    if !scanner.Scan() {
		panic("Scan returned false")
	}
	return scanner.Text()
}
func readInt() int {
	num, _ := strconv.Atoi(read())
	return num
}
func main() {
    scanner = bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    n := readInt()
    
    var no []int
    for i := 0 ; i < n; i++ {
        curr := readInt()
        
        no = append(no, curr)
    }
    
    loc := 1
    max := 0
    for i := 1 ; i < n; i++ {
        if no[i] >= no[i-1] {
            loc++
            continue
        }
        if max < loc {
            max = loc       
        }
        loc = 1
    }
    if max < loc {
        max = loc       
    }
    fmt.Println(max)
}
