package main

import (
	"belajar_goLang/database"
	_ "belajar_goLang/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
