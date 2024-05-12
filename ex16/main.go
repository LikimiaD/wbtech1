package main

import "fmt"

// Base: https://youtu.be/WprjBK0p6rw?si=yNpr6OOFZlPLkDzJ

func swap(x, y *int) {
	*x = *x ^ *y //?  X XOR Y
	*y = *x ^ *y //? (X XOR Y) XOR Y = X XOR (Y XOR Y) = X XOR 0 = X
	*x = *x ^ *y
}

func quicksort(slice []int) {
	if len(slice) < 2 {
		return //? Элементы уже отсортированы
	}

	swapMarker := -1             //? Указатель на swap, середина массива в конце цикла
	pivot := slice[len(slice)-1] //? last element

	//? Перебор всех элементов
	for index := 0; index < len(slice); index++ {
		if slice[index] > pivot {
			continue //? Значение больше нашего опорного элемента, ничего не делаем
		} else {
			swapMarker++ //? Двигаем маркер для swap
			if index > swapMarker {
				swap(&slice[index], &slice[swapMarker])
				//? index > swapMarker ? swap : pass
			}
		}
	}
	quicksort(slice[:swapMarker])   //? Сортировка левой части
	quicksort(slice[swapMarker+1:]) //? Сортировка правой части
}

func main() {
	slice := []int{3, 2, 5, 0, 1, 8, 7, 6, 9, 4}
	quicksort(slice)
	fmt.Println(slice)
}
