package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func main() {
	msg = "Hello world!"

	wg.Add(2)
	go updateMessage("Hello, universe!")
	go updateMessage("Hello, cosmos!")
	wg.Wait()

	fmt.Println(msg)
}

//func main() {
//msg = "Hello world!"

//var mutex sync.Mutex

//wg.Add(2)
//go updateMessage("Hello, universe!", &mutex)
//go updateMessage("Hello, cosmos!", &mutex)
//wg.Wait()

//fmt.Println(msg)
//}
