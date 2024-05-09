package main

import "fmt"

type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

type NotFoundError struct {
	Message string
}

func (n *NotFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &ValidationError{"validation error"}
	}

	if id != "kayu" {
		return &NotFoundError{"data not found"}
	}

	return nil
}

func main() {
	err := SaveData("kayu", nil)
	if err != nil {
		// if validationErr, ok := err.(*ValidationError); ok {
		// 	fmt.Println("validation error:", validationErr.Error())
		// } else if notFoundErr, ok := err.(*NotFoundError); ok {
		// 	fmt.Println("not found error:", notFoundErr.Error())
		// } else {
		// 	fmt.Println("unknow error", err.Error())
		// }
		switch finalError := err.(type) {
		case *ValidationError:
			fmt.Println("validation error:", finalError.Error())
		case *NotFoundError:
			fmt.Println("not found error:", finalError.Error())
		default:
			fmt.Println("unknow error:", finalError.Error())
		}
	} else {
		fmt.Println("success")
	}
}
