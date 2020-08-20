package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	go send(even, odd, quit)

	// receive
	receive(even, odd, quit)

	fmt.Println("About to exit")
}

func send(even, odd chan <- int, quit chan <- bool) {
	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	close(quit)
}

func receive (even, odd <- chan int, quit <- chan bool) {
	for {
		select {
		case value := <-even:
			fmt.Println("From the even channel:", value)
		case value := <-odd:
			fmt.Println("From the odd channel:", value)
		case value, ok := <-quit:
			if !ok {
				fmt.Println("From comma ok", value, ok)
				return
			} else {
				fmt.Println("From comma ok", value)
			}
		}
	}
}