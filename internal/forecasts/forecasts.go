package forecasts

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Interval struct {
	timestamp int64
	textTime  string
}

func GenerateForecastResponse(name string, lat, lon float64) ForecastResponse {
	response := ForecastResponse{
		Cod:     "200",
		Message: 0,
		Cnt:     40,
		List:    getWeatherInfo(),
		City: CityInfo{
			Id:   rand.Intn(999999),
			Name: name,
			Coord: Coord{
				Lat: lat,
				Lon: lon,
			},
		},
	}

	return response
}
func getWeatherInfo() []WeatherInfo {
	var weatherInfo []WeatherInfo
	intervals := GetIntervals()
	for _, interval := range intervals {
		info := WeatherInfo{
			Dt:         interval.timestamp,
			Main:       MainInfo{Temp: randFloat(200, 300), FeelsLike: rand.Intn(300), TempMin: randFloat(200, 300), TempMax: randFloat(200, 300), Pressure: rand.Intn(1200), SeaLevel: rand.Intn(1200), GrndLevel: rand.Intn(1500), Humidity: rand.Intn(100), TempKf: randFloat(-10, 10)},
			Weather:    []Weather{{Id: 802, Main: "Clouds", Description: "scattered clouds", Icon: "03n"}},
			Clouds:     Clouds{All: rand.Intn(50)},
			Wind:       Wind{Speed: randFloat(0, 10), Deg: rand.Intn(360), Gust: randFloat(0, 10)},
			Visibility: 10000,
			Pop:        0,
			Sys:        Sys{Pod: "n"},
			DtTxt:      interval.textTime,
		}

		weatherInfo = append(weatherInfo, info)
	}
	return weatherInfo
}

func GetIntervals() []Interval {
	// Текущее время
	now := time.Now()

	// Определяем ближайшее будущее время из интервалов по 3 часа
	nextInterval := getNextInterval(now)

	// Форматируем время в текстовом формате
	formattedTime := nextInterval.Format("2006-01-02 15:04:05")
	fmt.Println("Ближайшее будущее время:", formattedTime)

	// Создаем массив с временными метками на трое суток
	timeIntervals := generateIntervals(nextInterval)

	return timeIntervals
}

func getNextInterval(t time.Time) time.Time {
	// Берем только часы и минуты
	hour := t.Hour()
	// minute := t.Minute()

	// Определяем следующий интервал
	nextHour := (hour/3 + 1) * 3

	// Если следующий час больше 24, то переходим на следующий день
	if nextHour >= 24 {
		nextHour = 0
		t = t.Add(24 * time.Hour)
	}

	// Возвращаем новое время с нулями в минутах и секундах
	return time.Date(t.Year(), t.Month(), t.Day(), nextHour, 0, 0, 0, t.Location())
}

func generateIntervals(start time.Time) []Interval {
	var intervals []Interval
	currentTime := start

	// Генерируем 24 значения с интервалом в 3 часа
	for i := 0; i < 40; i++ {
		intervals = append(intervals, Interval{
			timestamp: int64(currentTime.Unix()),
			textTime:  currentTime.Format("2006-01-02 15:04:05"),
		})
		currentTime = currentTime.Add(3 * time.Hour)
	}

	return intervals
}

func randFloat(min, max float64) float64 {
	result := min + rand.Float64()*(max-min)
	limitedResult, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64)
	return limitedResult
}
func RandInit() {
	// Инициализируем генератор случайных чисел с текущего времени
	rand.New(rand.NewSource(time.Now().UnixNano()))
}
