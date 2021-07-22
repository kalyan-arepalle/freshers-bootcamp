package main

import (
	"fmt"
	"sync"
	"time"
)

func withdraw(balance *int, amount int, m *sync.Mutex,w *sync.WaitGroup){
	defer w.Done()
	m.Lock()
	if amount > *balance {
		fmt.Println("Insufficient Balance")
	} else {
		*balance -= amount
		fmt.Println("Balance after withdraw :",*balance)
	}
	m.Unlock()
}

func deposit(balance *int, amount int, m *sync.Mutex,w *sync.WaitGroup){
	defer w.Done()
	m.Lock()
	*balance += amount
	fmt.Println("Balance after deposit :",*balance)
	m.Unlock()
}

func main(){
	var m = &sync.Mutex{}
	var w = &sync.WaitGroup{}
	balance := 500

	for i:=0; i< 5; i++ {
		w.Add(1)
		go deposit(&balance,100,m,w)
	}
	time.Sleep(10*time.Microsecond)
	for i:=0; i< 5; i++ {
		w.Add(1)
		go withdraw(&balance,200,m,w)
	}
	w.Wait()
	fmt.Println(balance)
}