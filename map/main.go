package main

import "fmt"

func main() {

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"black": "#000000",
	// }

	// or

	// var colors map[string]string
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	colors := make(map[int]string)
	colors[10] = "#ffffff"
	fmt.Println(colors)

	//Delete map key value by passing map name and key name
	delete(colors, 10)
	fmt.Println("after delete", colors)
}
