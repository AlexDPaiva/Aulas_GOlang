package main

import "fmt"

func main() {
	// ponto de ebulição em kelvin 373,20
	var ebuK = 373.20
	var kelvin = 273

	fmt.Println("O valor de ebulição em celcios é ", ebuK - float64(kelvin))
}