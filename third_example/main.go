package main

import (
	"fmt"
)

func createGorutines(num int) <-chan int {

	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 1; i <= num; i++ {
			ch <- i
		}
	}()

	return ch
}

func main() {
	var number int
	fmt.Scan(&number)

	channel := createGorutines(number)
	sum := 0
	for val := range channel {
		sum += val
	}
	fmt.Printf("Sum = %d\n", sum)
}
