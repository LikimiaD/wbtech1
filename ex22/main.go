package main

import (
	"fmt"
	"math/big"
)

func main() {
	//? Инициализация базовых переменных
	//? .SetPrec(64) устанавливает точность в 64 бита
	a := big.NewFloat(13.132).SetPrec(64)
	b := big.NewFloat(12.132).SetPrec(64)

	sum := new(big.Float).Add(a, b) //? Сложения двух переменных
	fmt.Println("a + b -> ", sum)   //? Вывод

	sub := new(big.Float).Sub(a, b) //? Вычитание из переменной a переменную b
	fmt.Println("a - b -> ", sub)   //? Вывод

	div := new(big.Float).Quo(a, b) //? Умножение переменных
	fmt.Println("a / b -> ", div)   //? Вывод

	mul := new(big.Float).Mul(a, b) //? Деление переменной a на переменную b
	fmt.Println("a * b -> ", mul)   //? Вывод
}
