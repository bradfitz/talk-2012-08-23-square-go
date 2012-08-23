package main

import "fmt"

func f(left chan<- int, right <-chan int) {
	left <- 1 + <-right
}

func main() {
	const n = 100000
	leftmost := make(chan int)
	left, right := leftmost, leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func() { right <- 1 }()
	fmt.Println(<-leftmost)
}
