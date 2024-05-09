package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "kayu",
		"address": "kuningan",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)
	fmt.Println(person["salah"])

	books := map[string]string{
		"title":  "buku go blok",
		"author": "kayu",
		"salah":  "ups",
	}

	fmt.Println(books)

	delete(books, "salah")

	fmt.Println(books)
}
