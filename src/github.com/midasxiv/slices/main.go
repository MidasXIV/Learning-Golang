package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learning About Slices")

	grades := []int{34, 45, 23} // slice syntax
	fmt.Printf("Grades %v\n", grades)
	fmt.Printf("len of grades: %v\n", len(grades))
	fmt.Printf("capacity of grades: %v\n", cap(grades))

	gradesCopy := grades // reference and not a deep copy unlike arrays
	gradesCopy[1] = 50
	fmt.Printf("Grades  %v\n", grades)
	fmt.Printf("Grades Copy %v\n", gradesCopy)

	a := []int{10, 20, 30, 40, 50, 60}
	b := a[:]   // slice of all elements
	c := a[3:]  // slice of 4th element to end
	d := a[:3]  // slice of first 4 elements
	e := a[3:5] // slice of first 4 elements

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
