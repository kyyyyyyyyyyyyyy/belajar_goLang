package main

import (
	"belajar_goLang/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("kayu")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version)
	// fmt.Println(helper.sayGoodBye)
}
