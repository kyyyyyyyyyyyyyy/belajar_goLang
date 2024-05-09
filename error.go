package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("tidak boleh di gunakan")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := pembagian(100, 0)
	if err == nil {
		fmt.Println("hasil", hasil)
	} else {
		fmt.Println("error", err.Error())
	}
}
