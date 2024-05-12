package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup  //? Инициализация группы ожидания
	var total atomic.Int64 //? Инициализация атомарного счетчика
	total.Store(0)         //? Счетчик начинает с 0-ля

	for i := 0; i < 5; i++ { //? 5 горутин добавляют по 10.000 в total
		wg.Add(1) //? Добавление операции в группу ожидания
		go func() {
			defer wg.Done() //? Информируем группу ожидания, что операция завершилась
			for i := 0; i < 10000; i++ {
				total.Add(1)
			}
		}()
	}

	wg.Wait()                          //? Ожидание завершения всех горутин
	fmt.Println("total", total.Load()) //? Вывод значения после конкурентного добавления
}
