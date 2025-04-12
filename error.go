package main

import (
	"fmt"
	"strconv"
)

func main() {
	_, err := strconv.Atoi("123u")
	if err != nil {
		fmt.Println("Error", err)
	}
}
