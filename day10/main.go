package main

import (
	"fmt"
	"myapp/services"
	"myapp/utils"
)

func main() {
	// kita panggil disini

	// nama package utils memanggil function sayhello
	utils.SayHello("Aris")

	result := services.Add(10, 5)
	fmt.Println("Results :", result)

}
