// Запуск горутины и ожидание её завершения
// Задача: Напишите функцию, которая запускает горутину,
// выполняющую fmt.Println("Hello from goroutine!"),
// и использует sync.WaitGroup для ожидания её завершения.

package main

import (
	"fmt"
	"sync"
)

func printHello() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {

		defer wg.Done()

		fmt.Println("Hello from goroutine!")

	}()

	wg.Wait()
}

func main() {
	printHello()
}
