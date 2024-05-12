package main

import (
	"fmt"
	"strings"
)

func isEachSymbolUnique(line string) bool {
	line = strings.ToLower(line)        //? Делаем чтобы все символы были в нижнем регистре
	searcher := make(map[rune]struct{}) //? Map, где хранится информация о символах, которые проверили

	for _, sym := range line {
		if _, ok := searcher[sym]; ok {
			return false //? Нашли повторение
		} else {
			searcher[sym] = struct{}{} //? Записали новый символ
		}
	}
	return true //? Не нашли повторения
}

// ? Чтобы не писать много раз fmt.Printf()
func check(line string) {
	fmt.Printf("Checking line '%s', result -> %t\n", line, isEachSymbolUnique(line))
}

func main() {
	check("abcd")
	check("abCdefAaf")
	check("aabcd")
}
