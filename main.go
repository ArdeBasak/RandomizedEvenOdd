package main

import (
	"EvenOdd/service"
	"fmt"
)

func main() {
	// Create an instance of the mock data provider
	dataProvider := &service.DataProviderImpl{}

	// Call ConsumeData with a single input
	result, err := service.ConsumeData(dataProvider, 5)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Print the result
	fmt.Printf("Result: %s\n", result)
}
