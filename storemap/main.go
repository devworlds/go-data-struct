package main

import "fmt"

func main() {
	store := NewStore()

	store.Set("user", "jao123")
	fmt.Printf("%s", store.Get("user"))
}
