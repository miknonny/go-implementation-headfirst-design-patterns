package main

func main() {
	weatherData := NewWeatherData()

	// on creating a NewCurrentCondition we subscibed to the weather data.
	condition1 := NewCurrentCondition(1, weatherData)
	weatherData.RegisterObserver(condition1)

	weatherData.SetMeasurement(23, 902, 200)

	weatherData.NotifyObservers()

	weatherData.RemoveObserver(condition1)

	weatherData.SetMeasurement(1, 1, 1)
	weatherData.NotifyObservers()

}
