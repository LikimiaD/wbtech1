package main

import (
	"fmt"
	"sync"
)

/*
Основная идея для демонстрации будет array из JavaScript
P.s. Map(Object) где ключ -> индекс
*/

/*
Если мы планируем часто добавлять, редактировать, то
нам нужно реализация через Mutex
*/
func manyAddAndEdit() map[int]int {
	var mutex sync.Mutex         //? Блокировка, для записи в Map
	var wg sync.WaitGroup        //? Группа ожидания
	obj := make(map[int]int)     //? Инициализация Map
	arr := [5]int{1, 2, 3, 4, 5} //? Инициализация массива
	for index, value := range arr {
		wg.Add(1) //? Добавляем в счетчик информацию о запуске функции
		go func(index, value int) {
			defer wg.Done() //? Информируем что программа закончила работу
			mutex.Lock()    //? Блокировка для записи\чтения из вне
			obj[index] = value
			mutex.Unlock() //? Разблокировка для записи\чтения из вне
		}(index, value)
	}
	wg.Wait() //? Ожидание пока все горутины закончат работу
	return obj
}

/*
Если мы планируем один раз записать и
множество раз читать, то
можно рассмотреть sync.Map
*/
func oneAddManyTake() *sync.Map {
	var obj sync.Map             //? Специальная конкурентная Map'a
	var wg sync.WaitGroup        //? Группа ожидания
	arr := [5]int{1, 2, 3, 4, 5} //? Инициализация массива
	for index, value := range arr {
		wg.Add(1) //? Добавляем в счетчик информацию о запуске функции
		go func(index, value int) {
			defer wg.Done()         //? Информируем что программа закончила работу
			obj.Store(index, value) //? Запись в конкурентную Map
		}(index, value)
	}
	wg.Wait() //? Ожидание пока все горутины закончат работу
	return &obj
}

func main() {
	fmt.Printf("mutex map -> %v\n", manyAddAndEdit())
	m := oneAddManyTake()
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true
	})
}
