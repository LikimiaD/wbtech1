package main

import "fmt"

func binarySearch(slice []int, searchItem int) (bool, int, int) {
	low := 0               //? Начало отрезка на котором поиск
	high := len(slice) - 1 //? Последний элемент массива
	isContains := false    //? Ответ, существует или нет

	for low <= high || !isContains {
		middle := (low + high) / 2 //? Отвечает за средний элемент в массиве
		guess := slice[middle]     //? Получение среднего элемента
		if guess == searchItem {   //? Нашли искомое
			isContains = true
			return isContains, guess, middle
		} else if guess > searchItem { //? Среднее число больше искомого, оно в левой части среза
			high = middle - 1
		} else { //? Среднее число меньше искомого, оно в правой части среза
			low = middle + 1
		}
	}
	return isContains, 0, -1 //? Элемент не найден
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	found, value, index := binarySearch(slice, 4)
	fmt.Printf("Trying to find 4 -> found: %t, value: %d, index: %d\n", found, value, index)
}
