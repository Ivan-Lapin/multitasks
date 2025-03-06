// Fan-In (объединение данных из нескольких каналов)
// Задача: Напишите функцию, которая объединяет два входных канала в один выходной.

package main

import (
	"fmt"
	"sync"
	"time"
)

func merge(cs ...<-chan int) <-chan int {

	var wg sync.WaitGroup
	out := make(chan int)

	for _, in := range cs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range in {
				out <- n
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in1 := make(chan int)
	in2 := make(chan int)

	// Запускаем горутины для отправки данных в каналы
	go func() {
		for i := 1; i <= 5; i++ {
			in1 <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(in1)
	}()

	go func() {
		for i := 6; i <= 10; i++ {
			in2 <- i
			time.Sleep(150 * time.Millisecond)
		}
		close(in2)
	}()

	// Объединяем каналы
	out := merge(in1, in2)

	for n := range out {
		fmt.Printf("%d ", n)
	}

}
