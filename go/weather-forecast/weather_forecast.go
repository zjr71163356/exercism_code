// Package weather is a package about weather forecast
// provide function to forecast weather.
package weather

// CurrentCondition represent current weather condition  of the  city using string.
var CurrentCondition string

// CurrentLocation is the name of the city which is specified.
var CurrentLocation string

// Forecast returns a string value represent the weather condition of specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
