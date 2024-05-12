package main

import (
	"fmt"
	"sort"
)

func main() {
	temperatureSlice := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} //? Инициализация массива
	sort.Slice(temperatureSlice, func(i, j int) bool {
		return temperatureSlice[i] < temperatureSlice[j]
	}) //? Сортировка значения по возрастанию

	ans := make(map[int][]float64) //? Map для ответа

	for _, value := range temperatureSlice {
		bucket := int(value/10) * 10             //? Получение диапазон кратный 10
		ans[bucket] = append(ans[bucket], value) //? Записываем результат в правильный диапазон
	}

	fmt.Printf("%v\n", ans)
}
