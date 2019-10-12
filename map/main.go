package main

import "fmt"

func main() {
	// The first string is for a key, and the second is for a value
	colors := map[string]string{
		"red":   "#FF0000",
		"black": "#000000",
	}

	/* This is a different way to create a map */
	// colors := make(map[string]string)

	colors["white"] = "#ffffff"
	delete(colors, "white")
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}
