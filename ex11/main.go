package main

import "fmt"

func findIntersection(firstSlice, secondSlice []int) []int {
	if len(firstSlice) == 0 || len(secondSlice) == 0 {
		return make([]int, 0)
	} //? Одного из массива не существует

	searcher := make(map[int]struct{}) //? Карта для проверки
	for _, value := range firstSlice {
		searcher[value] = struct{}{} //? Запись уникальных значений
	}

	ans := make([]int, 0)
	for _, value := range secondSlice {
		if _, ok := searcher[value]; ok { //? Проверка наличия значения из первого массива
			ans = append(ans, value) //? Если есть, то записываем в ответ
		}
	}
	return ans
}

func main() {
	firstSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	secondSlice := []int{2, 4, 6, 8, 10}
	fmt.Printf("Intersection between first and second slice -> %v\n", findIntersection(firstSlice, secondSlice))
}
