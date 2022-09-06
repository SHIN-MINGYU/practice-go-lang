package main

import (
	"fmt"

	"github.com/learngo/dict"
)

func main2() {
	dictionary := dict.Dictionary{"first" : "first word"}

	// definition, err := dictionary.Search("second")
	// if(err != nil){
	// 	fmt.Println(err)
	// } else{
	// 	fmt.Println(definition)
	// }

	err := dictionary.Add("hello", "Greeting")

	if(err != nil){
		fmt.Println(err)
	}
	err2 := dictionary.Update("hello", "Greeting for Viewers")
	if (err2 != nil){
		fmt.Println(err2)
	}
	definition, err3 := dictionary.Search("hello")

	fmt.Println(definition)
	if (err3 != nil){
		fmt.Println(err3)
	}

	dictionary.Delete("hello")

	definition2, err4 := dictionary.Search("hello")
	if (err4 != nil){
		fmt.Println(err4)
	}
	fmt.Println(definition2)
}