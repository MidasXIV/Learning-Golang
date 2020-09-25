package main

import (
	"fmt"
	"reflect"
)

// a tag is specified after the type of the feild
// it has no meaning in Golanf but it is used
// by internal logic

// it is standard to use space between key-value pairs
// and the value to be in double quotes

type Animal struct {
	Name   string `required:"true" max:"100"`
	Origin string
}

func main() {
	fmt.Println("Learning about Struct Tag in Golang")
	tag := reflect.TypeOf(Animal{})
	field, _ := tag.FieldByName("Name")
	fmt.Println(field)     //{Name  string required:"true" max:"100" 0 [0] false}
	fmt.Println(field.Tag) //required:"true" max:"100"

	// this is passed to internal logic to handle the tags
}
