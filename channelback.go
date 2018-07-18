package main

import (
	"fmt"
	"time"
)

func add(t int) chan int {
	ch := make(chan int, 1)
	go func() {
		for i := 0; i < t; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func test4() {
	ticker := time.NewTicker(5 * time.Second)
	ticker1 := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	chh := add(10)
	end := make(chan struct{})
	for {

		select {
		case e, ok := <-chh:
			if ok == false {
				return
			}
			fmt.Println(e, "xxxxxxxxxxxxxx") // if ch not empty, time.After will nerver exec
			fmt.Println("sleep one seconds ...")
			go func() {
				time.Sleep(1 * time.Second)
				end <- struct{}{}
			}()
			select {
			case <-ticker1.C:
				fmt.Println("内部 timeout")
			case <-end:
			}
			fmt.Println("sleep one seconds end...")
		default: // forbid block
		}
		select {
		case <-ticker.C:
			fmt.Println("timeout")
			return
		default: // forbid block
		}
	}
}
func main() {
	test4()
}
