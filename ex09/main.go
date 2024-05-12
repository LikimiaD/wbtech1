package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func addNewValue(ctx context.Context, out chan<- int64) {
	defer close(out) //? Закрытие канала после завершения(сигнала отмены)

	rand.New(rand.NewSource(time.Now().UnixNano())) //? Указываем Seed

	for {
		select {
		case <-ctx.Done(): //? Ожидание сигнала завершения-отмены
			fmt.Println("addNewValue(): stop adding values")
			return
		default:
			value := rand.Int63n(100) //? Генерация случайного числа в int64
			fmt.Printf("addNewValue(): new X -> %d\n", value)
			out <- value //? Отправка числа в канал записи
			time.Sleep(1 * time.Second)
		}
	}
}

func doubleValue(ctx context.Context, in <-chan int64, out chan<- int64) {
	defer close(out) //? Закрытие канала после завершения(сигнала отмены)

	for {
		select {
		case value, ok := <-in: //? Получение результата из канала записи
			if !ok { //? Дополнительная проверка на наличие значения после получения
				fmt.Println("doubleValue(): input channel closed")
				return
			}
			doubled := value * 2
			fmt.Printf("doubleValue(): doubled X -> %d\n", doubled)
			out <- doubled //? Запись модифицированного числа во второй канал записи
		case <-ctx.Done(): //? Ожидание сигнала завершения-отмены
			fmt.Println("doubleValue(): stop doubled values")
			return
		}
	}
}

func printValue(in <-chan int64) {
	for value := range in { //? Чтение из модифицированного канала записи(второй канал)
		fmt.Printf("printValue(): %d\n", value)
	}
	fmt.Println("printValue(): done printing all values")
}

func main() {
	firstChannel := make(chan int64)                                         //? Первый канал для отправки новых значений
	secondChannel := make(chan int64)                                        //? Второй канал с информацией о модифицированных символах
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //? Информируем что через 10 с. надо закончить работу
	defer cancel()

	go addNewValue(ctx, firstChannel)
	go doubleValue(ctx, firstChannel, secondChannel)
	go printValue(secondChannel)

	time.Sleep(11 * time.Second) //? Ожидание срабатывания timeout'a
}
