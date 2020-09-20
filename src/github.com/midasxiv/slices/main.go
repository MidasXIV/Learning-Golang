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
	c := a[3:]  // slice of 4th element to end [inclusive]
	d := a[:3]  // slice of first 4 elements [ exclusive]
	e := a[3:5] // slice of first 4 elements

	a[3] = 23 // changes are reflected in b,c,d,e

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// using make to create a slice
	//<type>,<length>,<capacity>(not required)
	gg := make([]int, 3)
	fmt.Println(len(gg))
	fmt.Println(cap(gg))

	// unlike array slices don't need to have fixed sizes
	// hence length and capacity are different

	fmt.Printf("Slice :: %v | len :: %v | capacity :: %v\n", gg, len(gg), cap(gg))
	gg = append(gg, 1) // takes 2 arguments first is slice var and then the value
	fmt.Printf("Slice :: %v | len :: %v | capacity :: %v\n", gg, len(gg), cap(gg))
	gg = append(gg, 2, 3, 4)
	fmt.Printf("Slice :: %v | len :: %v | capacity :: %v\n", gg, len(gg), cap(gg))

	// the capacity incerease with each append as go creates a new slice each time
	// an operation is performed hence not good;

	// How to concatenate slices
	gc := []int{23, 32, 2323, 32232}
	gg = append(gg, gc...) // similiar to js spread operator
	fmt.Printf("Slice :: %v | len :: %v | capacity :: %v\n", gg, len(gg), cap(gg))
	//> Slice :: [0 0 0 1 2 3 4 23 32 2323 32232] | len :: 11 | capacity :: 12

	// Trim elements
	gg = gg[1:]         //remove first
	gg = gg[:len(gg)-1] //remove last
	fmt.Printf("Slice :: %v | len :: %v | capacity :: %v\n", gg, len(gg), cap(gg))
	// Slice :: [0 0 1 2 3 4 23 32 2323] | len :: 9 | capacity :: 11

	// Remove element from middle
	fmt.Println()
	fmt.Printf("GG :: %v | len :: %v | capacity :: %v\n", gg, len(gg), cap(gg))
	fmt.Printf("GC :: %v | len :: %v | capacity :: %v\n", gc, len(gc), cap(gc))

	fmt.Println("removing 1,2,3,4 from GG and assigning that to GC")
	gc = append(gg[:2], gg[6:]...)

	fmt.Printf("GG :: %v | len :: %v | capacity :: %v\n", gg, len(gg), cap(gg))
	fmt.Printf("GC :: %v | len :: %v | capacity :: %v\n", gc, len(gc), cap(gc))

	// you'll notice GG is now mutated; there's no safe way to do this yet

}
