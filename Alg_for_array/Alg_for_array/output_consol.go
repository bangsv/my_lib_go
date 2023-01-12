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

func Ellipsis(str string) {
	fmt.Print(str)
	for {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
		fmt.Print(".")
		time.Sleep(1 * time.Second)
		fmt.Print(".")
		time.Sleep(1 * time.Second)
		fmt.Print(".")
		fmt.Print("\b\b\b   \b\b\b\b")
	}
}
