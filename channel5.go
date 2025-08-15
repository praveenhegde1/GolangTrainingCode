package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	cities := []string{"EUR", "GBP", "CAD"}

	for _, city := range cities {
		url := fmt.Sprintf("https://apilayer.net/api/live?access_key=3b34bbee75bbd24586d0153c38fbbe8f&currencies=%s&source=USD&format=1", city)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Failed to get weather data for %s: %v\n", city, err)
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Failed to read response body for %s: %v\n", city, err)
			continue
		}

		var weather WeatherData
		err = json.Unmarshal(body, &weather)
		if err != nil {
			fmt.Printf("Failed to unmarshal weather data for %s: %v\n", city, err)
			continue
		}

		fmt.Printf("City: %s\n", weather.Name)
		fmt.Printf("Temperature: %.2f°C\n", weather.Main.Temp-273.15)
		fmt.Println()
	}
}
