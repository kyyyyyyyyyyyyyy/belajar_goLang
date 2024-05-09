package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("perulangan ke", counter)
	// 	counter ++
	// }

	for counter := 3; counter <= 10; counter++ {
		fmt.Println("perulangan ke", counter)
	}

	fmt.Println("done")

	names := []string{"muhammad", "rizky"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
