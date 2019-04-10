package main

import(
    "fmt"
    "time"
)

func main(){
    testArckermann()
}

func testArckermann() {
    fmt.Printf("start run\n")
    var m uint64
    var n uint64
    fmt.Scanln(&m)
    fmt.Scanln(&n)
    fmt.Printf("m: %d, n: %d\n", m, n)
    resultChan := make(chan uint64)
    go func() {
        resultChan <- arckermann(m, n)
    }()
    startTime := time.Now()
    loop:
    for {
        select{
        case result := <-resultChan:
            fmt.Printf("arckermann(%d, %d) = %d\n", m, n, arckermann(m, n), result)
            break loop
        case <- time.After(time.Second*2):
            fmt.Printf("runing waited times: %ds\n", time.Now().Sub(startTime)/time.Second)
        }
    }
}

func arckermann(m uint64, n uint64) uint64 {
    for m!=0 {
        if n == 0 {
            n = 1
        } else {
            n = arckermann(m, n-1)
        }
        m -= 1
    }
    return n+1
}
