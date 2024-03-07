package main

import "fmt"

func convertCelsiusToFahrenheit(c float64) float64{
	return(c * 9/5) + 32
}

func convertFahrenheitToCelsius(f float64) float64{
	return (f - 32) * 5/9
}

func main(){
	var choice int
	var temp float64


	fmt.Println("Welcome to Temperature Converter!")
	fmt.Println("Select an option (1 or 2)...")
	fmt.Println("1. Convert Celsius to Fahrenheit")
	fmt.Println("2. Convert Fahrenheit to Celsius")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Print("Enter temperature in Celsius: ")
		fmt.Scanln(&temp)
		fmt.Printf("Temperature in Fahrenheit: %.2f\n", convertCelsiusToFahrenheit(temp))
	case 2:
		fmt.Print("Enter temperature in Fahrenheit: ")
		fmt.Scanln(&temp)
		fmt.Printf("Temperature in Celsius: %.2f\n", convertFahrenheitToCelsius(temp))
	default:
		fmt.Println("Invalid choice. Please enter 1 or 2")
	}

}