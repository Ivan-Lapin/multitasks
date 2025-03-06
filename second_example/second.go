// Запуск нескольких горутин и ожидание их завершения
// Задача: Напишите программу, которая запускает 5 горутин,
// каждая из которых печатает свой номер (от 1 до 5),
// и использует sync.WaitGroup для их синхронизации(нужно подождать их выполнения).

package main

import (
	"fmt"
	"sync"
)

func startGorutines(number int) {
	var wg sync.WaitGroup

	for i := 1; i <= number; i++ {

		wg.Add(1)

		go func() {

			defer wg.Done()

			fmt.Printf("Index of a gorutine: %d\n", i)

		}()
	}

	wg.Wait()
}

func main() {
	var count int
	fmt.Scan(&count)

	startGorutines(count)
}
