package main

// TODO: implement the function convertCelsiusToFahrenheit
func convertCelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

// TODO: implement the function convertFahrenheitToCelsius
func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	// TODO: call the function convertCelsiusToFahrenheit
	convertCelsiusToFahrenheit(0)	// 32
	convertCelsiusToFahrenheit(100) // 212
	convertCelsiusToFahrenheit(20)  // 68
	// TODO: call the function convertFahrenheitToCelsius
	convertFahrenheitToCelsius(32)  // 0
	convertFahrenheitToCelsius(212) // 100
	convertFahrenheitToCelsius(68)  // 20
}
