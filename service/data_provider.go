package service

// DataProvider is an interface for fetching and retrieving data.
type DataProvider interface {
	GetRandomNumber(id int) (int, error)
}
