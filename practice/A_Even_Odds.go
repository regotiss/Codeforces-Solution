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
func readInt64() int64 {
	num, _ := strconv.ParseInt(read(), 10, 64)
	return num
}

func main() {
    scanner = bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    n := readInt64()
    k := readInt64()
    hf := (n+1)/2
    
    ans := k*2 - 1
    if k > hf {
        ans = (k - hf)*2
    }
    
    fmt.Println(ans)
}
