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
	var wg sync.WaitGroup //? Ожидающая группа для синхронизации горутин

	slice := []int64{2, 4, 6, 8, 10} //? Инициализация массива

	for _, value := range slice {
		wg.Add(1) //? Добавление информации о запуске функции
		go func(number int64) {
			defer wg.Done() //? Уведомляем что функция закончила работу
			fmt.Printf("Base -> %d | Squared -> %d\n", number, numberSquared(number))
		}(value)
	}
	wg.Wait() //? Ожидание пока все функции закончат работу
}
