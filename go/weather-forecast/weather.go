// Package weather comment1
package weather

// CurrentCondition comment2
var CurrentCondition string
// CurrentLocation comment3
var CurrentLocation string

// Forecast comment4
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
