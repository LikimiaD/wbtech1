package main

import "fmt"

func getSet(stringSlice []string) []string {
	if len(stringSlice) == 0 { //? Проверка, что массив существует
		return make([]string, 0)
	}

	searcher := make(map[string]struct{}) //? Запись уникальных значений
	for _, value := range stringSlice {
		searcher[value] = struct{}{}
	}

	ans := make([]string, 0, len(searcher)) //? Создаем срез с выделанным буфером под все уникальные значения
	for value := range searcher {
		ans = append(ans, value)
	}
	return ans
}

func main() {
	stringSlice := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Printf("Intersection -> %v\n", getSet(stringSlice))
}
