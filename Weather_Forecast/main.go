package main

import (
	"encoding/json"
	"html/template"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
)

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	City    string
	Weather []Weather `json:"weather"`
	Main    Main      `json:"main"`
	Wind    Wind      `json:"wind"`
	Clouds  Clouds    `json:"clouds"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Main struct {
	Temp float64 `json:"temp"`
}

type Weather struct {
	Description string `json:"description"`
}

type Wind struct {
	Speed float64 `json:"speed"`
}

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"ApiKey"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}
	return c, nil
}

func query(city string) (Welcome, error) {
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return Welcome{}, err
	}
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiConfig.OpenWeatherMapApiKey)
	if err != nil {
		return Welcome{}, err
	}

	defer resp.Body.Close()

	var d Welcome
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return Welcome{}, err
	}
	return d, nil

}

func main() {
	log.Println("The server was started!")
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			//Dynamic URLs(?) Anyway, getting the city from URL
			city := strings.SplitN(r.URL.Path, "/", -1)[1]
			//Get-request
			data, err := query(city)
			if err != nil {
				http.Error(w, err.Error(),
					http.StatusInternalServerError)

				return
			}
			t, _ := template.ParseFiles("index.html")

			//I have created an extra struct-field to transmit
			//to the HTML-template
			data.City = strings.SplitN(r.URL.Path, "/", -1)[1]
			//Convert from float64 Kelvins to truncated Celsius
			data.Main.Temp -= 273
			data.Main.Temp = math.Trunc(data.Main.Temp*10) / 10

			//Executing index.html with data structure transmitted
			t.Execute(w, data)
			//CLI Input :D
			log.Println(data)
		})
	http.ListenAndServe(":8080", nil)
}
