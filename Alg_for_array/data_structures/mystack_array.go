package Alg_for_array

import (
	"fmt"
	"os"
)

type stack struct {
	count   int
	element []int
}

func (elemets *stack) push_back(data int) {
	elemets.element = append(elemets.element, data)
	elemets.count += 1
}

func (elemets *stack) pop_back() int {
	var data int
	if elemets.count == 0 {
		panic("Stack empty!")
	}
	data = elemets.element[elemets.count-1]
	elemets.count = elemets.count - 1
	elemets.element = elemets.element[:elemets.count]
	return data
}

func (elemets *stack) empty() int {
	if elemets.count == 0 {
		return 0
	}
	return 1
}

func (elemets *stack) top() int {
	var data int

	if elemets.count == 0 {
		panic("Stack empty!")
	}

	data = elemets.element[0]
	elemets.count = elemets.count - 1
	elemets.element = elemets.element[elemets.count:]

	return data
}

func testing(data *stack) bool {
	var chose int

	print("\nВыберите дейсвтие со стеком:\n 1 - Заполнить стек\n 2 - Получить последний входящий элемент\n 3 - Получить первый элемент стека с начало\n 4 - Проверить пустой ли стек\n5 - выход")
	fmt.Fscan(os.Stdin, &chose)
	if chose == 1 {
		var element int
		fmt.Println("\nВведите элемент: ")
		fmt.Fscan(os.Stdin, &element)
		data.push_back(element)
		testing(data)
	} else if chose == 2 {
		fmt.Println(data.pop_back())
		testing(data)
	} else if chose == 3 {
		fmt.Println(data.top())
		testing(data)
	} else if chose == 4 {
		fmt.Println(data.empty())
		testing(data)
	} else if chose == 5 {
		return true
	}
	return false
}
