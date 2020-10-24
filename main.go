package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)


type WeatherResponse struct {
		Coord struct {
			Lon float64
			Lat float64
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
			Main        string
			Description string
			Icon        string
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
			All float64
		}
		Name string
	}

func main() {
	var (
		city  string
		units string
		appid string
		debug bool
		lang  string
	)

	appid := "shhh..its a secret"
	
	flag.StringVar(&lang, "l", "en", "Language")
	flag.StringVar(&city, "p", "London", "Place")
	flag.StringVar(&units, "m", "metric", "Units in Metric, US, UK, etc")
	flag.BoolVar(&debug, "d", false, "Debug (default: False)")

	flag.Parse()

	url := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&units=" + units + "&appid=" + appid + "&lang=" + lang

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

	

	var f WeatherResponse
	err = json.Unmarshal(data, &f)
	if err != nil {
		fmt.Printf("Error unmarshall %s", err)
	}

	fmt.Printf("Current weather at %s, %s is:\n", f.Name, f.Sys.Country)
	fmt.Printf("Latitude: %.2f, Longitude: %.2f \n", f.Coord.Lat, f.Coord.Lon)
	fmt.Printf("Current temperature is %.2f%s \n", f.Main.Temp, degree)
	fmt.Printf("Feels like %.2f%s \n", f.Main.Feels_like, degree)
	fmt.Printf("Description: %s \n", f.Weather[0].Description)
	fmt.Printf("Wind Speed %.2f%s \n", f.Wind.Speed, speed)
	fmt.Printf("Clouds %.2f \n", f.Clouds.All)
	fmt.Printf("Visibility %.2f \n", f.Visibility)
	sr := time.Unix(int64(f.Sys.Sunrise), 0)
	ss := time.Unix(int64(f.Sys.Sunset), 0)
	fmt.Printf("Current time is %s\n", time.Now().Format("3:05PM"))
	fmt.Printf("Sunrise is at %d:%d:%d \n", sr.Hour(), sr.Minute(), sr.Second())
	fmt.Printf("Sunset is at %d:%d:%d \n", ss.Hour(), ss.Minute(), ss.Second())

}
