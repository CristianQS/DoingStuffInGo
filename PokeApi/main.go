package main

import (
	"PokeApi/domain"
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
	var locations domain.Location
	req,_ := http.NewRequest("GET","https://pokeapi.co/api/v2/location/canalave-city",nil)
	resp,_ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &locations)

	csvFile, err := os.Create("locations.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer  csvFile.Close()

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
