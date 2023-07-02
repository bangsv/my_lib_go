package Alg_for_array

import "fmt"

type stack struct {
	data int
	next *stack
}

func (data *stack) push_back(element int) *stack {

	node := &stack{data: element, next: data}
	return node

}

func (data *stack) pop_back() (int, *stack) {

	elemet := data.data
	return elemet, data.next

}

func (data *stack) empty() bool {

	if data == nil {
		return false
	}

	return true
}

func testing_method(mystack *stack) {

	// Testing push
	fmt.Println(mystack)
	mystack = mystack.push_back(12)
	fmt.Println(mystack)
	mystack = mystack.push_back(13)
	fmt.Println(mystack)
	mystack = mystack.push_back(14563)
	fmt.Println(mystack)

	// Testing pop
	print("\n")
	qwe, mystack := mystack.pop_back()
	fmt.Println(qwe)
	qwr, mystack := mystack.pop_back()
	fmt.Println(qwr)
	fmt.Println(mystack.empty())
	rew, mystack := mystack.pop_back()
	fmt.Println(rew)
	fmt.Println(mystack.empty())
	// rer, mystack := mystack.pop_back()
	// fmt.Println(rer)

}
