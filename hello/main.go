package main

import "fmt"

func main() {
	var name string
	fmt.Println("Write your name: ")

	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("Hello,", name)
}
