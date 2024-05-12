package main

import (
	"fmt"
	"time"
)

func sleep(dur time.Duration, cancel <-chan struct{}) {
	select {
	case <-time.After(dur):
		//? Время задержки истекло
	case <-cancel:
		//? Получен сигнал отмены
		return
	}
}

func main() {
	fmt.Println("start func with some logic...")
	cancel := make(chan struct{}) //? Создание канала отмены
	go func() {
		time.Sleep(2 * time.Second) //? Отмена вызывается через 2 секунды
		fmt.Println("cancel func...")
		close(cancel)
	}()
	sleep(5*time.Second, cancel)
	fmt.Println("func done his work...")
}
