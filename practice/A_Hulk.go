package main

import (
    "os"
    "bufio"
    "fmt"
    "strconv"
    "strings"
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
    
    f := []string{"I hate", "I love"}
    var fls []string
    
    for i := 0 ; i < n; i++ {
        fls =  append(fls, f[i % 2])
    }
    
    fmt.Println(strings.Join(fls, " that "), "it")
}
