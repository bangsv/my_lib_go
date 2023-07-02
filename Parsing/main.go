package main

import (
	"encoding/json"
	"fmt"
	"modul/modul/proccesing_string"
	"os"
	"strings"

	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

type Phone struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

var mydb []Phone

func main() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"https://avrora24.ru/catalog/telefoniya_i_svyaz/mobilnye_telefony/filter/city-belgorod/?PAGEN_1=3"},
		ParseFunc: parsing,
	}).Start()
}

func WriteJSONToFile(data []Phone, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return err
	}

	fmt.Println("JSON data has been written to", filename)
	return nil
}

func parsing(g *geziyor.Geziyor, r *client.Response) {

	data := r.HTMLDoc.Find("div.item-title").Text()
	price := r.HTMLDoc.Find("div.my-1").Text()

	processing_data(data, price)

	print_struct(mydb)

	err := WriteJSONToFile(mydb, "phones.json")
	if err != nil {
		fmt.Println("Error writing JSON:", err)
	}
}

func output_info(info []string) {
	// Выводим заполненный массив
	for _, str := range info {
		fmt.Println(str)
	}
}

func (data Phone) print() {
	fmt.Println("Name :", data.Name, "---> Price: ", data.Price)
}

func print_struct(data []Phone) {
	for item := range data {
		data[item].print()
	}
}

func processing_data(data string, price string) {

	data_array := strings.Split(data, " ")
	price_array := strings.Split(price, " ")

	new_data := proccesing_string.Split_str_marks_phone(data_array)
	new_price := proccesing_string.Split_str_price_phone(price_array)

	for index, darray := range new_data {

		d := darray
		p := new_price[index]
		db := Phone{
			Name:  d,
			Price: p,
		}
		mydb = append(mydb, db)
	}
}
