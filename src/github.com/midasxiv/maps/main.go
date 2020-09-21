package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learning maps")
	// map<keyword>[<type of key>]<type of value>
	carMiles := map[string]int{
		"Honda":  12345,
		"Nissan": 23222,
		"Toyota": 23443,
		"Dodge":  55643,
	}

	fmt.Println(carMiles)

	// if you don't have the value before hand use make to create map
	carModel := make(map[string]string)
	carModel = map[string]string{
		"Honda":  "civic",
		"Nissan": "sentra",
		"Toyota": "camry",
		"Dodge":  "dart"}

	fmt.Println(carModel["Dodge"])

	// adding data to map
	carModel["Mazda"] = "6"
	fmt.Println(carModel)

	// The order of the map is not consisistent
	// it'll be rearranged

	// remove from a map
	delete(carModel, "Dodge")
	fmt.Println(carModel)

	fmt.Println(carModel["typo"]) // gives zero value of "value type" hence ""
	fmt.Println(carMiles["typo"]) // gives zero value of "value type" hence 0

	mile, validEntry := carMiles["typo"]
	fmt.Println(mile, validEntry) // will return zero value and validEntry will give false

	mile, validEntry = carMiles["Dodge"]
	fmt.Println(mile, validEntry) //55643 true

	// to check only if a key is present

	_, ok := carMiles["typo"] // underscore skips the value and can't be used
	fmt.Println(ok)

	// length of maps
	fmt.Println(len(carModel))

}
