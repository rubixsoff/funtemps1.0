package conv

import "math"

// FahrenheitToCelsius converts a temperature in Fahrenheit to Celsius.
func FahrenheitToCelsius(f float64) float64 {
	return round((f-32)*5/9, 2)
}

// CelsiusToFahrenheit converts a temperature in Celsius to Fahrenheit.
func CelsiusToFahrenheit(c float64) float64 {
	return round((c*9/5 + 32), 2)
}

// KelvinToCelsius converts a temperature in Kelvin to Celsius.
func KelvinToCelsius(k float64) float64 {
	return round((k - 273.15), 2)
}

// CelsiusToKelvin converts a temperature in Celsius to Kelvin.
func CelsiusToKelvin(c float64) float64 {
	return round((c + 273.15), 2)
}

func KelvinToFahrenheit(k float64) float64 {
	return round((k-273.15)*9/5+32, 2)
}

// FahrenheitToKelvin converts a temperature in Fahrenheit to Kelvin.
func FahrenheitToKelvin(f float64) float64 {
	return round((f-32)*5/9+273.15, 2)
}

func round(f float64, n int) float64 {
	scale := math.Pow(10, float64(n))
	return math.Round(f*scale) / scale
}
