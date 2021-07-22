package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var sum = 0

func student(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	sum += rand.Intn(5)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 200; i++ {
		wg.Add(1)
		go student( &wg)
	}
	wg.Wait()

	fmt.Println(sum)
	avg := float64(sum)/200
	fmt.Println(avg)
}