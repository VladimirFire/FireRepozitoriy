package main

import (
	"fmt"
)

func main() {
    var N,N2,N3 uint64
	N=0
    N2=1
    N3=N+N2
    fmt.Println("100 ЧИСЕЛ ФИбОНАЧИ:")
    fmt.Println(N)
    fmt.Println(N2)
for i := 1; i <= 100; i++ {
    
  fmt.Println(N3)
    N=N2
    N2=N3
    N3=N+N2
    
}

}
