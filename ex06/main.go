package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var stopSignal bool
var mutex sync.Mutex

// ! Функция завершает работу после выполнения всех инструкций
func defaultComplete() {
	fmt.Println("Goroutine complete his work and stopped :)")
}

// ! Функция завершает работу по истечение времени работы через timeout
func stopWithTimeout(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopWithTimeout(): timeout called, existing...")
			return
		default:
			fmt.Println("stopWithTimeout(): it's time to work (╯°□°)╯︵ ┻━┻")
		}
		time.Sleep(1 * time.Second)
	}
}

// ? Хоть Timeout является оберткой над Deadline,
// ? но чисто теоретически это два разных подхода ;)
// ! Функция завершает работу по истечение времени работы через deadline
func stopWithDeadline(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopWithDeadline(): deadline called, existing...")
			return
		default:
			fmt.Println("stopWithDeadline(): it's time to work (╯°□°)╯︵ ┻━┻")
		}
		time.Sleep(1 * time.Second)
	}
}

// ! Функция завершает работу при получении сигнала о завершении
func stopWithSelfCloseChannel(ch <-chan struct{}) {
	for {
		select {
		case <-ch:
			fmt.Println("stopWithSelfCloseChannel(): grided value from our close channel, existing...")
			return
		default:
			fmt.Println("stopWithSelfCloseChannel(): it's time to work (╯°□°)╯︵ ┻━┻")
		}
		time.Sleep(1 * time.Second)
	}
}

// ! Функция завершает работу при выполнении условия завершения
func stopWithSignal() {
	for {
		mutex.Lock()
		if stopSignal {
			fmt.Println("stopWithSignal(): Stop signal received, existing...")
			mutex.Unlock()
			return
		}
		mutex.Unlock()
		fmt.Println("stopWithSignal(): it's time to work (╯°□°)╯︵ ┻━┻")
		time.Sleep(1 * time.Second)
	}
}

// ! Функция завершает работу паникой и восстанавливается
func stopWithRandomPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("stopWithRandomPanic(): panic recovered:", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopWithRandomPanic(): created random panic ¯\\_(ツ)_/¯")
			panic("some error")
		default:
			fmt.Println("stopWithRandomPanic(): it's time to work (╯°□°)╯︵ ┻━┻")
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go defaultComplete()

	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelTimeout()
	go stopWithTimeout(ctxTimeout)
	time.Sleep(6 * time.Second)

	ctxDeadline, cancelDeadline := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancelDeadline()
	go stopWithDeadline(ctxDeadline)
	time.Sleep(11 * time.Second)

	ch := make(chan struct{}, 1)
	go stopWithSelfCloseChannel(ch)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- struct{}{}
	}()
	time.Sleep(3 * time.Second)

	go stopWithSignal()
	go func() {
		time.Sleep(5 * time.Second)
		mutex.Lock()
		stopSignal = true
		mutex.Unlock()
	}()
	time.Sleep(6 * time.Second)

	go stopWithRandomPanic()
	time.Sleep(6 * time.Second)
}
