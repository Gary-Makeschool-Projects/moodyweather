package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// weather structure
type weather struct {
	// City
	City string `json:"name"`
	// Temperature structure
	Temp m `json:"main"`
}
type m struct {
	// Temperature
	Temperature float32 `json:"temp"`
	// FeelsLike
	FeelsLike float32 `json:"feels_like"`
	// TempMin
	TempMin float32 `json:"temp_min"`
	// TempMax
	TempMax float32 `json:"temp_max"`
	// Pressure
	Pressure float32 `json:"pressure"`
	// Humidity
	Humidity float32 `json:"humidity"`
}

func main() {
	// Load the .env file in the current directory
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// api key
	key := os.Getenv("KEY")
	fmt.Println(key)

	// base url
	base := "https://api.openweathermap.org/data/2.5/weather?q=%s&units=imperial&appid=%s"

	api := fmt.Sprintf(base, "London", key)
	// ready api url
	fmt.Println(api)

	// request to api
	resp, err := http.Get(api)
	if err != nil {
		log.Fatalln(err)
	}
	// response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// weather structure
	var t weather
	er := json.Unmarshal(data, &t)
	if er != nil {
		panic(er)
	}
	feelsLike := t.Temp.FeelsLike
	fmt.Println(int(feelsLike))
}
