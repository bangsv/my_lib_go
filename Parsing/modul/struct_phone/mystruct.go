package structphone

import "fmt"

type Phone struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

func Print_struct(data []Phone) {
	for item := range data {
		data[item].Print()
	}
}

func (data Phone) Print() {
	fmt.Println("Name :", data.Name, "---> Price: ", data.Price)
}
