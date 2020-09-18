package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learning About Arrays")

	grades := [3]int{34, 45, 23}
	fmt.Printf("Grades %v\n", grades)

	// no need to provide length when initilzing the array
	abcs := [...]string{"a", "b", "c"}
	fmt.Printf("abcs %v\n", abcs)

	var students [3]string
	fmt.Printf("Students %v\n", students)
	students[1] = "sai"
	fmt.Printf("Students %v\n", students)
	students[2] = "kai"
	fmt.Printf("Students %v\n", students)
	students[0] = "jai"
	fmt.Printf("Students %v\n", students)
	fmt.Printf("Number of Students: %v\n", len(students))

	// Deep copy
	studentsCopy := students                       // deep copy
	studentsCopy[1] = "xxx"                        // original is not mutated
	fmt.Printf("Students %v\n", students)          // Students [jai sai kai]
	fmt.Printf("Students Copy %v\n", studentsCopy) //Students Copy [jai xxx kai]

	// Pointer
	studentsPointer := &students                         // pointer
	studentsPointer[1] = "xxx"                           // original is mutated
	fmt.Printf("Students %v\n", students)                // Students [jai sai kai]
	fmt.Printf("Students Pointer %v\n", studentsPointer) //Students Copy [jai xxx kai]

}
