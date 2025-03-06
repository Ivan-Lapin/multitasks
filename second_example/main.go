package main

import (
	"fmt"
	"sync"
)

func startGorutines(number int) {
	var wg sync.WaitGroup

	for i := 1; i <= number; i++ {

		wg.Add(1)

		go func(n int) {

			defer wg.Done()

			fmt.Printf("Index of a gorutine: %d\n", i)

		}(number)
	}

	wg.Wait()
}

func main() {
	var count int
	fmt.Scan(&count)

	startGorutines(count)
}
