package main

import "fmt"

func main() {

	// Basic Calculator 

	var a, b float64
	var op string

	fmt.Print("Masukkan angka pertama: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan operator (+, -, *, /): ")
	fmt.Scan(&op)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scan(&b)

	switch op {
	case "+":
		fmt.Printf("Hasil: %.2f\n", a+b)
	case "-":
		fmt.Printf("Hasil: %.2f\n", a-b)
	case "*":
		fmt.Printf("Hasil: %.2f\n", a*b)
	case "/":
		if b != 0 {
			fmt.Printf("Hasil: %.2f\n", a/b)
		} else {
			fmt.Println("Error: Pembagian dengan nol bro.")
		}
	default:
		fmt.Println("Operator gak valid.")
	}

	// Temperature Calculator

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
