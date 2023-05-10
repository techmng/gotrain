package main

import (
	"fmt"
	"github.com/techmng/gotrain/pricing"
)

func main() {
	var shareName string
	fmt.Print("Enter Stock name: ")
	_, err := fmt.Scanf("%s", &shareName)
	if err != nil {
		return
	}
	message := fmt.Sprintf("price of %s share  is %f", shareName, pricing.Price(shareName))
	fmt.Println(message)
}
