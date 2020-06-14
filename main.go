package main

import "fmt"

func main() {

	i := 0
	for i < 3 {
		i = i + 1
		fmt.Printf("Python style of looping %d\n", i)
	}
	for j := 0; j < 3; j++ {
		fmt.Printf("Initialization and iteration %d\n", j)
	}

}
