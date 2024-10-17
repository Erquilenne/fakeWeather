package main

import (
	"encoding/json"
	forecasts "fakeweather/internal/forecasts"
	"fmt"
	"net/http"
	"strconv"
)

type GeoResponse struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

type ForecastResponse struct {
	Temperature float64 `json:"temp"`
	Condition   string  `json:"condition"`
}

type City struct {
	Name string
	Lat  float64
	Lon  float64
}

func main() {
	forecasts.RandInit()
	// resp := forecasts.GenerateResponse("Moscow")
	// fmt.Println(resp)
	http.HandleFunc("/geo/1.0/direct", geoHandler)
	http.HandleFunc("/data/2.5/forecast", forecastHandler)
	fmt.Println("Server is running on port 8082")
	http.ListenAndServe(":8082", nil)
}

func geoHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	respArr := make([]forecasts.CityResponse, 0)
	response := forecasts.GenerateCityResponse(q)
	respArr = append(respArr, response)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respArr)
}

func forecastHandler(w http.ResponseWriter, r *http.Request) {
	lat, errLat := strconv.ParseFloat(r.URL.Query().Get("lat"), 32)
	lon, errLon := strconv.ParseFloat(r.URL.Query().Get("lon"), 32)
	if errLat != nil || errLon != nil {
		http.Error(w, "Invalid latitude or longitude", http.StatusBadRequest)
		return
	}

	// Заранее подготовленный JSON (пример)
	forecastCities := map[string]string{
		"55.7504_37.6174":   "Moscow",
		"59.9387_30.3162":   "Saint Petersburg",
		"51.5073_-0.1276":   "London",
		"48.8588_2.3200":    "Paris",
		"40.7127_-74.0060":  "New York",
		"35.6828_139.7594":  "Tokyo",
		"-33.8698_151.2082": "Sydney",
		"52.5170_13.3888":   "Berlin",
		"41.8933_12.4829":   "Rome",
		"40.4167_-3.7035":   "Madrid",
		"42.3554_-71.0605":  "Boston",
		"38.8950_-77.0365":  "Washington",
		"34.0536_-118.2427": "Los Angeles",
		"31.2322_121.4692":  "Shanghai",
		"37.5666_126.9782":  "Seoul",
		"48.1371_11.5753":   "Munich",
		"52.3727_4.8936":    "Amsterdam",
		"50.4500_30.5241":   "Kyiv",
		"52.2319_21.0067":   "Warsaw",
	}

	key := fmt.Sprintf("%.4f_%.4f", lat, lon)
	fmt.Println(key)
	if city, found := forecastCities[key]; found {
		response := forecasts.GenerateForecastResponse(city, lat, lon)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}
