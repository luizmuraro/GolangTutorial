package main

import (
	"time"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func cleanUp() {

	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup: ", r)
	}
}

func say(s string) {
	defer cleanUp()
	for i := 0; i < 3; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
		if i == 2 {
			panic("Oh dear, a 2!")
		}
	}
	
}

func main() {
	wg.Add(1)
	go say("ronaldo")
	wg.Add(1)
	go say("brilha muito no corinthians")
//	wg.Wait()

	wg.Add(1)
	go say("Hi")
	wg.Wait()

}