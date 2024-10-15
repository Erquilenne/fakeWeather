package forecasts

type CityResInfo struct {
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Name string  `json:"name"`
}

func GenerateCityResponse(name string) CityResponse {
	cities := map[string]CityResInfo{
		"Москва":          CityResInfo{Lat: 55.7504, Lon: 37.6174, Name: "Moscow"},
		"Санкт-Петербург": CityResInfo{Lat: 59.9387, Lon: 30.3162, Name: "Saint Petersburg"},
		"Лондон":          CityResInfo{Lat: 51.5073, Lon: -0.1276, Name: "London"},
		"Париж":           CityResInfo{Lat: 48.8588, Lon: 2.3200, Name: "Paris"},
		"Нью-Йорк":        CityResInfo{Lat: 40.7127, Lon: -74.0060, Name: "New York"},
		"Токио":           CityResInfo{Lat: 35.6828, Lon: 139.7594, Name: "Tokyo"},
		"Сидней":          CityResInfo{Lat: -33.8698, Lon: 151.2082, Name: "Sydney"},
		"Берлин":          CityResInfo{Lat: 52.5170, Lon: 13.3888, Name: "Berlin"},
		"Рим":             CityResInfo{Lat: 41.8933, Lon: 12.4829, Name: "Rome"},
		"Мадрид":          CityResInfo{Lat: 40.4167, Lon: -3.7035, Name: "Madrid"},
		"Бостон":          CityResInfo{Lat: 42.3554, Lon: -71.0605, Name: "Boston"},
		"Вашингтон":       CityResInfo{Lat: 38.8950, Lon: -77.0365, Name: "Washington"},
		"Лос-Анджелес":    CityResInfo{Lat: 34.0536, Lon: -118.2427, Name: "Los Angeles"},
		"Шанхай":          CityResInfo{Lat: 31.2322, Lon: 121.4692, Name: "Shanghai"},
		"Сеул":            CityResInfo{Lat: 37.5666, Lon: 126.9782, Name: "Seoul"},
		"Мюнхен":          CityResInfo{Lat: 48.1371, Lon: 11.5753, Name: "Munich"},
		"Амстердам":       CityResInfo{Lat: 52.3727, Lon: 4.8936, Name: "Amsterdam"},
		"Киев":            CityResInfo{Lat: 50.4500, Lon: 30.5241, Name: "Kyiv"},
		"Варшава":         CityResInfo{Lat: 52.2319, Lon: 21.0067, Name: "Warsaw"},
	}
	city := cities[name]
	response := CityResponse{
		Name:       city.Name,
		LocalNames: map[string]string{},
		Lat:        city.Lat,
		Lon:        city.Lon,
		Country:    "RU",
		State:      "Moscow",
	}

	return response
}
