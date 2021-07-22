package main

import (
	"encoding/json"
	"fmt"
)

func findFrequency(stringChannel chan string, m map[string]int, done chan string){
	str := <-stringChannel
	for _,c := range str{
		m[string(c)]++
	}
	done <- str
}

func main() {
	frequencyMap := make(map[string]int)
	input := []string{"quick", "brown", "fox", "lazy", "dog"}

	stringChannel := make(chan string, len(input))
	done := make(chan string, len(input))
	for i:=0;i<len(input);i++{
		go findFrequency(stringChannel, frequencyMap, done)
	}
	for i := 0; i < len(input); i++ {
		stringChannel <- input[i]
	}
	for i := 0; i < 5; i++ {
		<-done
	}

	stringJson,err := json.Marshal(frequencyMap)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(stringJson))
}
