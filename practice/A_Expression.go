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

func findMax(a, b, c int) int {
    res1 := a + b
    if res1 < (a * b) {
        res1 = a * b
    }
    res := res1 + c
    if res < (res1 * c) {
        res = res1 * c
    }
    return res
}
func main() {
    scanner = bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    a := readInt()
    b := readInt()
    c := readInt()
    
    res1 := findMax(a, b, c)
    res2 := findMax(c, b, a)
    
    if res1 < res2 {
        res1  = res2   
    }
    fmt.Println(res1)
}
