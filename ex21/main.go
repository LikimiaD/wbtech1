package main

import "fmt"

type Animal interface {
	//? Интерфейс, который ожидает система
	makeSound()
}

type CatAdapter struct {
	//? CatAdapter адаптер для кошек
	cat Cat
}

func (c CatAdapter) makeSound() {
	c.cat.meow()
}

type DogAdapter struct {
	//? DogAdapter адаптер для собак
	dog Dog
}

func (d DogAdapter) makeSound() {
	d.dog.woof()
}

type Cat struct {
	//? Cat структура кошки со своим методом
	Breed string
	Name  string
	Age   int
}

func (c Cat) meow() {
	fmt.Printf("A cat named %s, breed %s and age %d year says: 'meow'\n", c.Name, c.Breed, c.Age)
}

type Dog struct {
	//? Dog структура собаки со своим методом
	Name string
	Task string
	Age  int
}

func (d Dog) woof() {
	fmt.Printf("A dog named %s and age %d year has task '%s', but says: 'woof'", d.Name, d.Age, d.Task)
}

func main() {
	cat := Cat{
		Breed: "Siamese",
		Name:  "Mira",
		Age:   1,
	} //? Инициализация кота
	dog := Dog{
		Name: "Kira",
		Task: "Keep the owners happy",
		Age:  2,
	} //? Инициализация собаки}

	var animal Animal

	//? Проверка адаптер

	animal = CatAdapter{cat: cat}
	animal.makeSound()

	animal = DogAdapter{dog: dog}
	animal.makeSound()
}
