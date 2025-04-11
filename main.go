package main

import "fmt"

func main() {
	var temp float64
	var unit string

	fmt.Print("Masukkan suhu (contoh: 30): ")
	fmt.Scan(&temp)
	fmt.Print("Masukkan satuan asal (C, F, K): ")
	fmt.Scan(&unit)

	switch unit {
	case "C", "c":
		fmt.Printf("Fahrenheit: %.2f°F\n", (temp*9/5)+32)
		fmt.Printf("Kelvin: %.2fK\n", temp+273.15)
	case "F", "f":
		fmt.Printf("Celsius: %.2f°C\n", (temp-32)*5/9)
		fmt.Printf("Kelvin: %.2fK\n", ((temp-32)*5/9)+273.15)
	case "K", "k":
		fmt.Printf("Celsius: %.2f°C\n", temp-273.15)
		fmt.Printf("Fahrenheit: %.2f°F\n", ((temp-273.15)*9/5)+32)
	default:
		fmt.Println("Satuan gak dikenal.")
	}
}
