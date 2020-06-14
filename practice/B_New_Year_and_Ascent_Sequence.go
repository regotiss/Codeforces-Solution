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
func readFloat() float64 {
    num, _ := strconv.ParseFloat(read(), 64)
    return num
}
func main() {
    scanner = bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    t := readInt()

    var minv, maxv []float64
    for i := 0; i < t; i++ {
        l := readInt()
        
        min := 1e9
        max := -1e9
        asc := false
        for j := 0; j < l; j++ {
            v := readFloat()
            
            if v > min {
                asc = true
            }
            if v < min {
                min = v
            }
            if v > max {
                max = v
            }
        }
        if asc {
            min = -1e9
            max = 1e9
        }
        minv = append(minv, min)
        maxv = append(maxv, max)
    }
    sort.Float64s(maxv)

    var cnt uint64
    
    for i := 0; i < t; i++ {
        ub := sort.Search(len(maxv), func(j int) bool { return maxv[j] > minv[i] })
        cnt += uint64(t - ub)
    } 
    fmt.Println(cnt)
}
