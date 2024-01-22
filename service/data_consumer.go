package service

// ConsumeData is a function that uses a DataProvider to fetch and process data.
func ConsumeData(provider DataProvider, id int) (string, error) {
	// Use GetRandomNumber to get a random number between 0 and id
	randomNumber, err := provider.GetRandomNumber(id)
	if err != nil {
		return "", err
	}

	// Check whether the value is even or odd
	result := checkEvenOrOdd(randomNumber)

	// Return the result
	return result, nil
}

// checkEvenOrOdd checks whether the given value is even or odd.
func checkEvenOrOdd(value int) string {
	if value%2 == 0 {
		return "Even"
	}
	return "Odd"
}
