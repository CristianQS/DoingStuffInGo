package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
 	client := &http.Client{
 		Timeout: 10 * time.Second,
	}
	var locations Location
	req,_ := http.NewRequest("GET","https://pokeapi.co/api/v2/location/canalave-city",nil)
	resp,_ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &locations)

	csvFile, err := os.Create("locations.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer  csvFile.Close()

	fmt.Println(locations)
	writer := csv.NewWriter(csvFile)

	var row []string
	row = append(row, strconv.Itoa(locations.Id))
	row = append(row, locations.Name)
	row = append(row, locations.Region.Name)
	row = append(row, locations.Region.Url)
	for _, name := range locations.Names {
		row = append(row, name.Name)
	}
	for _, game := range locations.Games {
		row = append(row, strconv.Itoa(game.Id))
		row = append(row, game.Generation.Name)
		row = append(row, game.Generation.Url)
	}

	for _, area := range locations.Areas {
		row = append(row, area.Name)
		row = append(row, area.Url)
	}

	writer.Write(row)

	// remember to flush!
	writer.Flush()
}

type Location struct {
	Id int `json:"ìd"`
	Name string `json:"name"`
	Region Region `json:"region"`
	Names [] Name `json:"names"`
	Games [] Game `json:"game_indices"`
	Areas [] Area `json:"areas"`
}

type Region struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Name struct {
	Name string `json:"name"`
	Language Language `json:"language"`
}

type Language struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type Game struct {
	Id int `json:"game_index"`
	Generation Generation `json:"generation"`
}

type Generation struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type Area struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

