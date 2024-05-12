package main

import (
	"fmt"
	"strings"
)

func reverseWords(input, separator string) string {
	words := strings.Split(input, separator) //? Разделение строки на слова по сепаратору
	//? Делаем счетчик с начала и с конца, пока они не встретились делаем swap
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, separator) //? Соединение слов обратно в строку с тем же сепаратором
}

func main() {
	x := "snow dog sun"
	fmt.Printf("'%s'\n", reverseWords(x, " "))
}
