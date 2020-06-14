package main

import "fmt"

func main() {

	var a int = 1
	fmt.Printf("Declaration and Initialization of a %d\n", a)

	var a1 = 1
	fmt.Printf("Without explicit type declaration %d\n", a1)

	a2 := 1
	fmt.Printf("Short hand notation of declaration and initialization %d\n", a2)

	var z int
	fmt.Printf("Declaration and zero value initialization of z %d", z)

}
