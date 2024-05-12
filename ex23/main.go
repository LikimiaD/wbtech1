package main

import "fmt"

func removeElementSlice(slice []any, element any) []any {
	if len(slice) == 0 {
		return slice //? Возвращаем пустой слайс, если он уже пуст
	}

	for index, value := range slice {
		if value == element {
			//? Удаление элемента, соединение частей до и после найденного элемента
			return append(slice[:index], slice[index+1:]...)
		}
	}
	//? Элемент не найден, возвращаем исходный слайс
	return slice
}

func main() {
	x := []any{"1", "x", "y"}
	y := removeElementSlice(x, "1") //? Удаление первого элемента
	fmt.Println(y)                  //! ["x", "y"]

	z := []any{"a", "b", "c"}
	z = removeElementSlice(z, "c") //? Удаление последнего элемента
	fmt.Println(z)                 //! ["a", "b"]

	w := []any{"a", "b", "c"}
	w = removeElementSlice(w, "d") //? Попытка удалить несуществующий элемент
	fmt.Println(w)                 //! ["a", "b", "c"]

	empty := make([]any, 0)
	empty = removeElementSlice(empty, "anything") //? Попытка удалить элемент из пустого слайса
	fmt.Println(empty)                            //! []
}
