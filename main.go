package main

import (
	//	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	debug := false

	city := "London"
	units := "metric"
	appid := "103e41e1f9fdd4e18a872b70f4a1c251"
	url := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&units=" + units + "&appid=" + appid
	if debug {
		fmt.Println(url)
	}

	var degree string
	var speed string
	if units == "metric" {
		degree = "°C"
		speed = "km/h"
	} else {
		degree = "°F"
		speed = "m/h"
	}

	fmt.Println("Getting weather for " + city)

	res, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error fetching weather data: %s", err)
	}

	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)

	if debug {
		fmt.Println(string(data))
	}

	type WeatherResponse struct {
		Coord struct {
			lon float64
			lat float64
		}
		Main struct {
			Temp       float64
			Feels_like float64
			Temp_Min   float64
			Temp_Max   float64
			Humidity   float64
			Pressure   float64
		}
		Visibility float64
		Weather    []struct {
			Main string
			Desc string
			Icon string
		}

		Wind struct {
			Speed float64
			Deg   float64
		}
		Dt  float64
		Sys struct {
			Country string
			Sunrise float64
			Sunset  float64
		}
		Timezone float64
		Clouds   struct {
			all float64
		}
	}

	var f WeatherResponse
	err = json.Unmarshal(data, &f)
	if err != nil {
		fmt.Printf("Error unmarshall %s", err)
	}

	fmt.Printf("Current temperature is %.2f%s \n", f.Main.Temp, degree)
	fmt.Printf("Feels like %.2f%s \n", f.Main.Feels_like, degree)
	fmt.Printf("Wind Speed %.2f%s \n", f.Wind.Speed, speed)
	fmt.Printf("Clouds %.2f \n", f.Clouds.all)
	fmt.Printf("Visibility %.2f \n", f.Visibility)
	fmt.Printf("Sunrise is at %.2f \n", f.Sys.Sunrise)
	fmt.Printf("Sunset is at %.2f \n", f.Sys.Sunset)

}
