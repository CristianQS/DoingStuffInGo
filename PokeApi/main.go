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
	GetLocation(client, locations)

	CreateCsvFrom(locations)
}

func CreateCsvFrom(locations domain.Location) {
	csvFile, err := os.Create("locations.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	row := SerializeLocation(locations)
	writer.Write(row)
	writer.Flush()
}


func GetLocation(client *http.Client, locations domain.Location) {
	req, _ := http.NewRequest("GET", "https://pokeapi.co/api/v2/location/canalave-city", nil)
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &locations)
}

func SerializeLocation(locations domain.Location) (row [] string) {
	row = append(row, strconv.Itoa(locations.Id))
	row = append(row, locations.Name)
	row = SerializeRegion(locations.Region, row)
	for _, name := range locations.Names { row = SerializeName(row, name) }
	for _, game := range locations.Games { row = SerializeGame(row, game) }
	for _, area := range locations.Areas { row = SerializeArea(row, area) }
	return row
}

func SerializeRegion(region domain.Region, row []string) []string {
	row = append(row, region.Name)
	row = append(row, region.Url)
	return row
}
func SerializeName(row []string, name domain.Name) []string {
	row = append(row, name.Name)
	return row
}

func SerializeGame(row []string, game domain.Game) []string {
	row = append(row, strconv.Itoa(game.Id))
	row = append(row, game.Generation.Name)
	row = append(row, game.Generation.Url)
	return row
}
func SerializeArea(row []string, area domain.Area) []string {
	row = append(row, area.Name)
	row = append(row, area.Url)
	return row
}
