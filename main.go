package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

type Weather struct {
	Status WeatherCondition `json:"status"`
}

type WeatherCondition struct {
	Water int `json:"Water"`
	Wind  int `json:"Wind"`
}

func index(w http.ResponseWriter, r *http.Request) {

	// run html file
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile("weather.json")
	if err != nil {
		fmt.Print(err)
	}

	var obj Weather

	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	dataJson()

	fmt.Fprintln(w, "The Weather is :")

	fmt.Fprintln(w, "wind :", obj.Status.Wind, "kmph")
	fmt.Fprintln(w, "water :", obj.Status.Water, "m")
	condition(w, obj)

}

func condition(w http.ResponseWriter, obj Weather) {
	if obj.Status.Wind <= 6 {
		fmt.Fprintln(w, "Wind :aman")
	}
	if obj.Status.Wind >= 7 && obj.Status.Wind <= 15 {
		fmt.Fprintln(w, "Wind :status siaga")
	}
	if obj.Status.Wind > 15 {
		fmt.Fprintln(w, "Wind :bahaya")
	}

	if obj.Status.Water <= 5 {
		fmt.Fprintln(w, "Water :aman")
	}
	if obj.Status.Water >= 6 && obj.Status.Water <= 8 {
		fmt.Fprintln(w, "Water :status siaga")
	}
	if obj.Status.Water > 8 {
		fmt.Fprintln(w, "Water :bahaya")
	}
}

func main() {
	address := "localhost:9090"
	http.HandleFunc("/", index)
	log.Printf("Your service is up and running at : " + address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("Error running service: ", err)
	}
}

func dataJson() {
	wind := rand.Intn(100)
	water := rand.Intn(100)
	data := Weather{
		Status: WeatherCondition{Wind: wind,
			Water: water},
	}

	file, _ := json.MarshalIndent(data, "", " ")

	//write the file
	_ = ioutil.WriteFile("weather.json", file, 0644)
}
