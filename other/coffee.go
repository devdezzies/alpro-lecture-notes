package main 

import "fmt"

func main(){
    var n, m, x, y int
    var total int
    
    fmt.Scan(&n, &m, &x, &y)
    // for (n / x != 0 && m / y != 0) {
    //     total++
    //     n = n - x
    //     m = m - y
    // }
    // fmt.Println(total)
    if n / x <= m / y {
        total = n / x
    } else {
        total = m / y
    }
    fmt.Println(total)
}