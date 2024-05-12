package main

import (
	"errors"
	"fmt"
	"unsafe"
)

var (
	ErrIndexOutOfRange = errors.New("index out of range") //? Собственная ошибка
)

func changeBit(value int64, index int) (int64, error) {
	totalBits := unsafe.Sizeof(value) * 8 //? Узнаем максимальный разбем в битах
	if index < 0 || index >= int(totalBits) {
		return value, ErrIndexOutOfRange //? Указано значения за пределами
	}

	return value ^ (1 << index), nil //? Замена бита
}

func main() {
	var x int64 = 5
	if newX, err := changeBit(x, 8); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Changed 8 bit -> %d\n", newX)
	}
}
