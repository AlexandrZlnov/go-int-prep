// Задача:
// Что выведет код

package main

import "fmt"

type ErrInvalidID struct{}

func (e *ErrInvalidID) Error() string {
	return "invalid id"
}
func getData(id int) (string, error) {
	var (
		err  *ErrInvalidID
		data string
	)
	if id == 0 {
		err = &ErrInvalidID{}
	} else {
		data = "some data"
	}
	return data, err
}
func main() {
	data, err := getData(12)
	if err != nil {
		fmt.Printf("ERROR %s\n", err) // ERROR invalid id
	} else {
		fmt.Println(data)
	}
}

// Объяснение:
// err внутри функции равен nil, но при возврате он упаковывается в интерфейс error,
// который уже не равен nil, и при печати fmt автоматически вызывает метод Error().
