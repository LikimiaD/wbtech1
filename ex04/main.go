package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

// ? Алфавит для генерации сообщения
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// ? Получение информации сколько worker'ов могут работать одновременно
func getMaxWorkers() int {
	var maxWorkers int
	fmt.Printf("Enter number of workers: ")
	_, err := fmt.Scanf("%d", &maxWorkers)
	for err != nil || maxWorkers <= 0 { //? Получен некорректный ответ
		fmt.Printf("Please enter a valid positive number of workers: ")
		_, err = fmt.Scanf("%d", &maxWorkers)
	}
	return maxWorkers
}

// ? Перед началом заполняем буферизированный канал ID worker'ов
func fillWorkers(pool chan<- int, maxWorkers int) {
	for id := 0; id < maxWorkers; id++ {
		pool <- id
	}
}

func RandStringBytes(n int) string {
	// ? Генерация случайного сообщения
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// ? Worker выполняет работу, генерирует сообщение и отправляет или получает сигнал об отмене
func worker(ctx context.Context, id int, ch chan<- string, pool chan int) {
	message := RandStringBytes(10)
	fmt.Printf("Worker #%d sends message: %s\n", id, message)
	select {
	case ch <- message:
		time.Sleep(time.Second)
		pool <- id //! Уведомляем что рабочий свободен
	case <-ctx.Done(): //! Уведомление об отмене операции
		return
	}
}

func main() {
	maxWorkers := getMaxWorkers()                           //? Получение информации о кол-во рабочих
	ch := make(chan string)                                 //? Канал для получения сообщений
	pool := make(chan int, maxWorkers)                      //? Pool worker'ов
	ctx, cancel := context.WithCancel(context.Background()) //? Контекст завершения
	defer cancel()

	//? Канал, слушающий когда пользователь завершит программу
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	fillWorkers(pool, maxWorkers) //? Присваивание уникального номера каждому worker'у

	go func() {
		//! Вечный цикл while
		for {
			select {
			case id := <-pool: //! Получен номер свободного рабочего
				go worker(ctx, id, ch, pool)
			case <-signalCh: //! Получен сигнал о завершении
				fmt.Println("Received Ctrl+C, exiting...")
				close(ch)   //! Закрываем канал сообщений
				close(pool) //! Закрываем канал рабочих
				return
			}
		}
	}()

	//? Выводи сообщения, которые получаем от worker'ов
	for message := range ch {
		fmt.Printf("Main received: %s\n", message)
	}
}
