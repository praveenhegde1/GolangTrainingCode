package main

import (
	"fmt"
	"sync"
)

type WeatherData struct {
	City        string
	Temperature float64
	Humidity    float64
}

func main() {

	weatherDataChan := make(chan WeatherData)

	var wg sync.WaitGroup

	cities := []string{"New York", "London", "Tokyo"}

	// Create a goroutine for each city
	for _, city := range cities {
		wg.Add(1)
		go func(city string) {
			defer wg.Done()

			// Make API call to collect weather data for the city
			// Replace this with your actual API call
			temperature, humidity := getWeatherData(city)

			// Create a WeatherData struct with the collected data
			weatherData := WeatherData{
				City:        city,
				Temperature: temperature,
				Humidity:    humidity,
			}

			// Send the weather data to the channel
			weatherDataChan <- weatherData
		}(city)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Close the channel to signal that no more data will be sent
	close(weatherDataChan)

	// Receive weather data from the channel and print it
	for weatherData := range weatherDataChan {
		fmt.Printf("City: %s\n", weatherData.City)
		fmt.Printf("Temperature: %.2f\n", weatherData.Temperature)
		fmt.Printf("Humidity: %.2f\n", weatherData.Humidity)
		fmt.Println()
	}
}

// Function to simulate API call and return random weather data
func getWeatherData(city string) (float64, float64) {
	// Replace this with your actual API call
	// Simulating by returning random values
	return 25.0, 60.0
}
