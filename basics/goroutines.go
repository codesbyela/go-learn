package main

import (
    "fmt"
    "time"
)

func worker(id int, ch chan int) {
    fmt.Printf("Worker %d started\n", id)
    time.Sleep(2 * time.Second)
    fmt.Printf("Worker %d finished\n", id)
    ch <- id
}

func main() {
    ch := make(chan int)
    for i := 1; i <= 5; i++ {
        go worker(i, ch)
    }
    for i := 1; i <= 5; i++ {
        id := <-ch
        fmt.Printf("Received result from worker %d\n", id)
    }
}