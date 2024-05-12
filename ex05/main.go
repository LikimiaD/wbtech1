package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"time"
)

// ? Алфавит для генерации сообщения
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// ? Получение информации сколько worker'ов могут работать одновременно
func getMaxWorkers() int {
	var maxWorkers int
	fmt.Printf("enter number of workers: ")
	_, err := fmt.Scanln(&maxWorkers)
	for err != nil || maxWorkers <= 0 { //? Получен некорректный ответ
		fmt.Printf("please enter a valid positive number of workers: ")
		_, err = fmt.Scanln(&maxWorkers)
	}
	return maxWorkers
}

// ? Получение информации через сколько секунд программа должна завершить работу
func getFuncDuration() int {
	var funcDuration int
	fmt.Printf("enter time of func duration(seconds): ")
	_, err := fmt.Scanln(&funcDuration)
	for err != nil || funcDuration <= 0 { //? Получен некорректный ответ
		fmt.Printf("please enter a valid positive number: ")
		_, err = fmt.Scanln(&funcDuration)
	}
	return funcDuration
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
	maxWorkers := getMaxWorkers()                                                                     //? Получение информации о кол-во рабочих
	funcDuration := getFuncDuration()                                                                 //? Получение информации через сколько секунд закончится работа
	ch := make(chan string)                                                                           //? Канал для получения сообщений
	pool := make(chan int, maxWorkers)                                                                //? Pool worker'ов
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(funcDuration)*time.Second) //? Контекст завершения через N секунд
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
				runtime.Gosched()
			case <-ctx.Done(): //! Получен сигнал о прохождении времени работы
				fmt.Println("received time expired, exiting...")
				close(ch)   //! Закрываем канал сообщений
				close(pool) //! Закрываем канал рабочих
				return
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
