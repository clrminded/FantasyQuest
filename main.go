package main

import (
	"fmt"
	"log"
)

func main() {
	items, err := getItemData()
	if err != nil {
		log.Fatalf("Error getting item data: %v", err)
	}

	data := string(items)
	fmt.Println(data)
}
