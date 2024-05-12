package main

import "fmt"

func reverse(input string) string {
	runes := []rune(input) //? Входное может состоять из Unicode символов
	//? Делаем счетчик с начала и с конца, пока они не встретились делаем swap
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes) //? Получили string -> возвращаем string
}

func main() {
	x := "WB Tech: level # 1 (Golang)"
	fmt.Printf("'%s'\n", reverse(x))
}
