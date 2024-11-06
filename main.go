package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


type GeoResponse struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

func main() {
	
	city := "Bordeaux" 
	apiKey := "c732a4f732342956ec521490b59a7dce"
	baseURL := "http://api.openweathermap.org/data/2.5/weather"

	url := fmt.Sprintf("%s?q=%s&appid=%s", baseURL, city, apiKey)


	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erreur lors de la requête HTTP: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture de la réponse: %v", err)
	}


	if resp.StatusCode != 200 {
		log.Fatalf("Erreur API, code de statut HTTP: %d\n", resp.StatusCode)
	}

	
	var geoData GeoResponse
	if err := json.Unmarshal(body, &geoData); err != nil {
		log.Fatalf("Erreur lors du parsing JSON: %v", err)
	}

	
	fmt.Printf("Ville: %s\n", geoData.Name)
	fmt.Printf("Latitude: %.6f\n", geoData.Latitude)
	fmt.Printf("Longitude: %.6f\n", geoData.Longitude)
}
