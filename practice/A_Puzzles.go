package main
 
import (
    "os"
    "bufio"
    "fmt"
    "strconv"
    "sort"
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
    n := readInt()
    m := readInt()
    
    var ps []int
    for i:= 0; i < m; i++ {
        ps = append(ps, readInt())
    }
    sort.Ints(ps)
    
    min := ps[n-1] - ps[0]
    for i:= 1; i <= (m-n); i++ {
        curr := ps[i+n-1] - ps[i]
        if min > curr {
            min = curr
        }
    }
    fmt.Println(min)
}
