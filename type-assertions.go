package main

import "fmt"

func main() {
	var i interface{} = "hello"

	S := i.(string)
	fmt.Println(S)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)

}
