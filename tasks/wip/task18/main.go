// Что выведет код и почему?

package main

import "fmt"

type MyError struct{}

func (MyError) Error() string {
	return "MyError!"
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	var err *MyError
	errorHandler(err)

	err = &MyError{}
	errorHandler(err)
}

// Ответ:
// Error: <nil>
// Error: MyError!

// В случае:
// var err *MyError
// errorHandler(err)
/*
	inrerface {
		type *MyError
		data nil
	}
*/

// В случае:
//err = &MyError{}
// errorHandler(err)
/*
	inrerface {
		type *MyError
		data &MyError{}
	}
*/
