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
    h := readInt()
    
    cnt := n
    for i := 0 ; i < n; i++ {
        ai := readInt()
        
        if ai > h {
            cnt++
        }
    }
    
    fmt.Println(cnt)
}
