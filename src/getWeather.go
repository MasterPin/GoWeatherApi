package main

import (
	"encoding/json"
	"fmt"
	"github.com/chris-dot-exe/AwesomeLog"
	"io/ioutil"
	"net/http"
	"os"
)

type HourlyUnits struct {
	Time          string `json:"time"`
	Temp          string `json:"temperature_2m"`
	ApperentTemp  string `json:"apparent_temperature"`
	Precipitation string `json:"precipitation"`
}

type WeatherData struct {
	Latitude             float64     `json:"latitude"`
	Longitude            float64     `json:"longitude"`
	GenerationTimeMS     float64     `json:"generationtime_ms"`
	UTCOffsetSeconds     int         `json:"utc_offset_seconds"`
	Timezone             string      `json:"timezone"`
	TimezoneAbbreviation string      `json:"timezone_abbreviation"`
	Elevation            float64     `json:"elevation"`
	HourlyUnits          HourlyUnits `json:"hourly_units"`
	Hourly               struct {
		Time          []string  `json:"time"`
		Temp          []float64 `json:"temperature_2m"`
		ApperentTemp  []float64 `json:"apparent_temperature"`
		Precipitation []float64 `json:"precipitation"`
	} `json:"hourly"`
}

func GetWeather() {
	resp, err := http.Get("https://api.open-meteo.com/v1/dwd-icon?latitude=49.0069&longitude=8.4037&hourly=temperature_2m,apparent_temperature,precipitation&timezone=Europe%2FBerlin")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var weatherData WeatherData

	err = json.Unmarshal(respData, &weatherData)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("%+v", weatherData)

}
