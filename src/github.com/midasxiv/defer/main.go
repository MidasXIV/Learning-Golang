package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Learning defer in Go")
	fmt.Println("Is used to push a function call after execution of the function")
	fmt.Println("but before the function is returned")
	fmt.Println("Often used to free-up resources")

	res, err := http.Get("http://www.google.com/robots.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer fmt.Println("This will be printed last as defer is LIFO")
	defer fmt.Println("Using defer to close Body resource")
	defer res.Body.Close() // will still work because executed after function
	robots, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)

	// Learning Panic
	// Panic is used inplace of "exceptions" or Error
	// we can define custom panic

	fmt.Println("start")
	defer fmt.Println("deffered")
	panic("custom panic")
	fmt.Println("end") // will not get printed out

	// order of execution is always function -> defer -> custom panic
	// start
	// deffered
	// Using defer to close Body resource
	// This will be printed last as defer is LIFO
	// panic: custom panic
}
