package main

import "fmt"

func main(){

	var i int = 4

	if i > 5  {
		fmt.Println(i)
	} else {
		fmt.Println("i is smaller than 4")
	}
}

//there are no terary expression in go. Sad! :( 