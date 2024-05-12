package main

import "fmt"

/*
Вообще изначально была версия:
func swap(x, y *int) {
	*x = *x + *y
	*y = *x - *y
	*x = *x - *y
}

Но тут в первой строчке `*x = *x + *y`
может быть переполнение :c
*/

/*
   0 XOR 0 = 0
   1 XOR 1 = 0
   0 XOR 1 = 1
   1 XOR 0 = 1
*/

func swap(x, y *int) {
	*x = *x ^ *y // <- X XOR Y
	*y = *x ^ *y // <- (X XOR Y) XOR Y = X XOR (Y XOR Y) = X XOR 0 = X
	*x = *x ^ *y
}

func main() {
	x := 5
	y := 2
	swap(&x, &y)

	fmt.Printf("Swapped values -> x: %d\ty: %d\n", x, y)
}
