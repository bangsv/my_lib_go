package Alg_for_array

import (
	"bufio"
	"fmt"
	"os"
)

func System_pause() {
	fmt.Println("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadByte()
}

func Out_array(array []int) {
	fmt.Print("{ ")
	for i := 0; i < len(array); i++ {
		if i == len(array)-1 {
			fmt.Print(array[i])
			continue
		}
		fmt.Print(array[i], ", ")
	}
	fmt.Print(" }\n")
}
