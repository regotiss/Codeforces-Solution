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

func readInt64() int64 {
	num, _ := strconv.ParseInt(read(), 10, 64)
	return num
}
 
func main() {
    scanner = bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    n := readInt64()
    n1 := n/2
    n2 := n1  + n%2
    
    ans := n1*(n1+1) - n2*n2
    
    fmt.Println(ans)
}
