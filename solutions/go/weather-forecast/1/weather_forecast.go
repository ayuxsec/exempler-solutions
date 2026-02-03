// Package weather provides you information
// about current weather condition.
package weather

var (
    // CurrentCondition represents the current weather condition.
	CurrentCondition string 
    // CurrentLocation represents the location for which the weather information applies.
	CurrentLocation  string 
)

// Forecast sets the current location and weather condition
// and returns a formatted weather report string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
