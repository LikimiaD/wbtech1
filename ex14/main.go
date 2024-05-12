package main

import (
	"fmt"
	"reflect"
)

func getTypeUsingFmt(value interface{}) string {
	return fmt.Sprintf("%T", value)
}

func getTypeUsingSwitch(value interface{}) string {
	t := reflect.TypeOf(value)
	switch v := value.(type) {
	case int:
		return fmt.Sprintf("int: %d", v)
	case float64:
		return fmt.Sprintf("float64: %f", v)
	case bool:
		return fmt.Sprintf("bool: %t", v)
	case string:
		return fmt.Sprintf("string: %s", v)
	default:
		// Проверяем, является ли тип каналом
		if t.Kind() == reflect.Chan {
			// Получаем тип элементов канала
			elemType := t.Elem()
			return fmt.Sprintf("Channel of %s", elemType)
		}
		return "Unknown type"
	}
}

func getTypeUsingReflect(value interface{}) (reflect.Type, reflect.Value) {
	return reflect.TypeOf(value), reflect.ValueOf(value)
}

func main() {
	var x interface{} = make(chan []int)

	// Определение типа через флаг
	fmtType := getTypeUsingFmt(x)
	fmt.Println("Using fmt.Sprintf:", fmtType)

	// Определение типа через switch оператор
	switchType := getTypeUsingSwitch(x)
	fmt.Println("Using type switch:", switchType)

	// Использование reflect.TypeOf и reflect.ValueOf для определения типа и значения
	reflectType, reflectValue := getTypeUsingReflect(x)
	fmt.Println("Using reflect:", reflectType, reflectValue)
}
