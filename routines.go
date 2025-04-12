package main

import "fmt"

func sayHello() {
	for i := 0; i < 1000000; i++ {
		fmt.Println(i)
	}
}

// func main() {
// 	go sayHello()
// 	fmt.Println("hello")
// }
