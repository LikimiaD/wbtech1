package main

import "fmt"

type Human struct {
	//? Human определяет общие атрибуты
	Name   string
	Gender string
}

func (h Human) Speak() {
	//? Функция структуры Human
	fmt.Printf("Hello, my name is %s and I am a %s.\n", h.Name, h.Gender)
}

type Action struct {
	//? Структура Action, которая содержит в себе реализацию структуры Human
	//! Демонстрирует композицию и делегирование в Go
	TheatricalProductionName string
	Director                 Human
	Actors                   []Human
}

func main() {
	//? Создание экземпляров структуры Human
	director := Human{Name: "Dasha", Gender: "female"}
	igor := Human{Name: "Igor", Gender: "male"}
	gleb := Human{Name: "Gleb", Gender: "male"}
	kate := Human{Name: "Kate", Gender: "female"}

	//? Инициализация Action от Human
	prod := Action{
		TheatricalProductionName: "Nutcracker",
		Director:                 director,
		Actors:                   []Human{igor, gleb, kate},
	}

	//! Демонстрируем возможность директора говорить, что подтверждает композицию
	prod.Director.Speak()

	fmt.Printf("Production information: %s\n", prod.TheatricalProductionName)
	fmt.Println("Actors in the production:")
	for _, actor := range prod.Actors {
		fmt.Printf("- %s\n", actor.Name)
		actor.Speak() //! Демонстрация, что каждый актёр также может говорить
	}
}
