package main

import "fmt"

func main() {

	var username string
	fmt.Println("Enter first name:")
	fmt.Scan(&username)
	fmt.Printf("print your username %v, and the type is %T", username, username)
	fmt.Print("test commit")
	
}
