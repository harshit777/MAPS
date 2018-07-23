package main

import (
	"fmt"
)

func main() {

	//var colors map[string]string   //empty map

	//colors := make(map[string]string)

	colors := map[string]string{ //keys and values
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex value for " + color + " is " + hex)
	}
}
