package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Learning defer - panic - recover")
	fmt.Println("Starting Program")
	panicker()
	fmt.Println("Endinging Program")
}

func panicker() {
	fmt.Println("About to panic")
	defer func() {
		fmt.Println("This is a nameless function")
		fmt.Println("Using recover will not save the current function from which panic is invoked")
		fmt.Println("But return control to the previous function in callstack")
		if err := recover(); err != nil {
			log.Println("Error", err)
		}
	}()
	panic("Panicking !!")
	fmt.Println("Because of panic we wont be able to see this line but we'll see the end line in main program")
}

// output
// Learning defer - panic - recover
// Starting Program
// About to panic
// This is a nameless function
// 2020/09/28 01:07:13 Error Panicking !!
// Endinging Program
