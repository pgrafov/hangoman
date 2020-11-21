package resources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func pickRandomKey(myMap map[string][]string) string {
	keys := make([]string, 0, len(myMap))
	for k := range myMap {
		keys = append(keys, k)
	}
	return keys[rand.Intn(len(myMap))]
}

func PickRandomCity() (string, string) {
	jsonFile, err := os.Open("resources/cities.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonData := make(map[string][]string)
	err = json.Unmarshal(byteValue, &jsonData)
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
	countryName := pickRandomKey(jsonData)
	return countryName, jsonData[countryName][rand.Intn(len(jsonData[countryName]))]
}
