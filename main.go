package main

import (
	"errors"
	"fmt"
	"os"
)

type NotFoundError struct {
	UserId int
}

func (err NotFoundError) Error() string {
	return fmt.Sprintf("user with id %d not found", err.UserId)
}

func SearchUser(id int) error {
	// some logic for search
	// ...
	// if not found
	var err NotFoundError
	err.UserId = id
	return err
}

func main() {
	// const id = 17
	// err := SearchUser(id)
	// if err != nil {
	// 	fmt.Println(err)
	// 	//type error checking
	// 	notFoundErr, ok := err.(NotFoundError)
	// 	if ok {
	// 		fmt.Println(notFoundErr.UserId)
	// 	}
	// }

	err := openFile("non-existing")
	if err != nil {
		fmt.Println(err.Error())
		// get internal error
		fmt.Println(errors.Unwrap(err))
	}
}

func openFile(filename string) error {
	if _, err := os.Open(filename); err != nil {
		return fmt.Errorf("error opening %s: %w", filename, err)
	}
	return nil
}
