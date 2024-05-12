package main

/*
? Проблематика вопроса: Утечка памяти
? Причины:
* Использование среза из большой строки `v` приводит к тому, что переменная `justString`
* хранит ссылку на часть исходного массива байт, что в свою очередь удерживает весь массив
* в памяти. Сборщик мусора не может освободить память, занятую `v`, пока `justString`
* активна, так как `justString` ссылается на начальную часть `v`.
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	//! justString = v[:100] -> утечка памяти, хранит ссылку на весь массив `v`
	justString = string([]byte(v[:100])) // Создаем новую строку из среза
}

/*
! Пример функции, создающей большую строку.
! Создает строку размером `size`, заполненную символами 'A'.
*/
func createHugeString(size int) string {
	hugeString := make([]byte, size)
	for i := range hugeString {
		hugeString[i] = 'A'
	}
	return string(hugeString)
}

/*
! Вызов someFunc() создает строку, которая после выхода из функции someFunc()
! не удерживает весь массив `v` в памяти, что предотвращает утечку памяти.
*/
func main() {
	someFunc()
}