package main

import (
	"fmt"
	"sync"
)

// ? Функция для вычисления квадрата числа
func numberSquared(number int64) int64 {
	return number * number
}

func main() {
	slice := []int64{2, 4, 6, 8, 10}   //? Инициализация массива
	ch := make(chan int64, len(slice)) //? Инициализация буферизированого канала для сохранения квадратов
	var wg sync.WaitGroup              //? Ожидающая группа для синхронизации горутин

	for _, value := range slice {
		wg.Add(1) //? Добавление информации о запуске функции
		go func(ch chan<- int64, number int64) {
			defer wg.Done()             //? Уведомляем что функция закончила работу
			ch <- numberSquared(number) //? Запись квадрата числа
		}(ch, value)
	}

	wg.Wait() //? Ожидание пока все функции закончат работу
	close(ch) //? Закрытие канала, так как все записали

	var sum int64 //? Инициализация переменной для вывода
	for n := range ch {
		//? Получаем квадрат числа из буферизированого канала
		sum += n
	}

	fmt.Printf("Sum -> %d\n", sum)
}
