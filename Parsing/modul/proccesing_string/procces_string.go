package proccesing_string

func Split_str_marks_phone(data []string) []string {
	var new_array_data []string
	var time_str string
	for _, item := range data {
		time_str += " " + item
		if item == "GB" {
			new_array_data = append(new_array_data, time_str)
			time_str = ""
		}
	}
	return new_array_data
}

func Split_str_price_phone(data []string) []string {
	var new_array_data []string
	var time_str string
	for _, item := range data {
		time_str += " " + item
		if item == "₽" {
			new_array_data = append(new_array_data, time_str)
			time_str = ""
		}
	}
	return new_array_data
}
