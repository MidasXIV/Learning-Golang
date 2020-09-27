package main

import "fmt"

func main() {
	fmt.Println("Leanring loops in Go")

	for i := 0; i < 5; i += 2 {
		fmt.Println(i)
	}

	// using two variables
	for i, j := 0, 10; i < 7 && j >= 3; i, j = i+1, j-1 {
		fmt.Println(i, j)
	}

	// leaving out the initializer and increment parts of a for loop
	// test syntax
	k := 0
	for k < 5 {
		fmt.Println(k)
		k++
	}

	// infinite syntax for-loops
	j := 0
	for {
		fmt.Println(j)
		j++

		if j == 5 {
			break
		}
	}

	// breaking out to labels used to exit branched loops
LoopStart:
	for i1 := 0; i1 < 5; i1++ {
		for j1 := 0; j1 < 5; j1++ {
			fmt.Printf("%v%v ", i1, j1)
			if i1 == 3 && j1 == 3 {
				// break           // case 1
				break LoopStart //case 2
			}
		}
		fmt.Println()
	}

	// case 1
	// 00 01 02 03 04
	// 10 11 12 13 14
	// 20 21 22 23 24
	// 30 31 32 33
	// 40 41 42 43 44

	//case 2
	// 00 01 02 03 04
	// 10 11 12 13 14
	// 20 21 22 23 24
	// 30 31 32 33

	fmt.Println()
	//looping over maps, slices, strings and arrays
	arr := [3]int{23, 34, 45}
	for key, value := range arr {
		fmt.Println(key, value)
	}
}
