package main

import "fmt"

func main() {
	/*
		Another way of declaring map

		var colors map[string]string

		it will produce GO ZERO VALUE which is map[]
	*/

	/*
		Another way of declaring map

		colors := map([string]string)
	*/

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#0ff000",
	}

	/* add value into existing map */
	colors["white"] = "#ffffff"

	/* delete value in map */
	delete(colors, "green")

	fmt.Println(colors)
}
