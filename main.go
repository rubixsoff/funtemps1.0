package main

import (
	"flag"
	"fmt"

	"rubixsoff/funtemps/funtemps1/conv"
)

// Definerer flag-variablene i hoved-"scope"

var fahr, celsius, kelvin float64
var out string

/*
 Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
er initialisert.
*/

func init() {

	// Implementerer parsing av flagg

	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
}

func main() {

	// Automatisk setter initial value som "0.0"
	var inputTemp float64

	// Definere flagg som jeg skal bruke i input fra kommandolinjen
	flag.Parse()

	// Switch avgjør hvilken input temperatur som er oppgitt
	switch {
	// Dersom -F flagget passerte, så blir verdien "fahr" tildelt til inputTemp. Samme gjelder -C og -K men da med deres egne verdier.
	case isFlagPassed("F"):
		inputTemp = fahr
	case isFlagPassed("C"):
		inputTemp = celsius
	case isFlagPassed("K"):
		inputTemp = kelvin
	//Feilmelding i output dersom man ikke skriver inn noe av det følgende ovenfor
	default:
		fmt.Println("No input temperature provided.")
		return
	}

	// Konverterer input temperatur til Celsius
	//Switch sjekker hvilket flagg som er oppgitt
	switch {
	case isFlagPassed("F"):
		celsius = conv.FahrenheitToCelsius(inputTemp)
	case isFlagPassed("K"):
		celsius = conv.KelvinToCelsius(inputTemp)
	}

	// Konverterer Celsius til output temperatur
	switch out {
	case "C":
		fmt.Printf("%.2f°C\n", celsius)
	case "F":
		fahr = conv.CelsiusToFahrenheit(celsius)
		fmt.Printf("%.2f°F\n", fahr)
	case "K":
		kelvin = conv.CelsiusToKelvin(celsius)
		fmt.Printf("%.2f°K\n", kelvin)
	default:
		fmt.Println("Invalid output temperature specified.")
		return
	}
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
