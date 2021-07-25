package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var sum = 0

func studentRating(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	sum += rand.Intn(5)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 200; i++ {
		wg.Add(1)
		go studentRating( &wg)
	}
	wg.Wait()

	avg := float64(sum)/200
	fmt.Println("The average rating for the teacher is : ",avg)
}