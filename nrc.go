package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// carMakes is a struct to hold the data from all makes
type carMakes struct {
	name   string
	makeID int
}

var baseURL = "https://vpic.nhtsa.dot.gov/api/vehicles/"
var makes []carMakes

func main() {
	fmt.Println("Starting NRC")
	GetAllMakes()
}

// GetAllMakes - gets every make returned from the API
func GetAllMakes() {
	url := baseURL + "GetAllMakes?format=json"
	var jsonData map[string]interface{}

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Get recall failed with error %s\n", err)
		return
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("error reading body %s\n", err)
	}

	err = json.Unmarshal(data, &jsonData)

	if err != nil {
		fmt.Printf("Received error unmarshalling %s\n", err)
	}

	for k, v := range jsonData {
		fmt.Printf("key: %s\n", k)
		if k == "Results" {
			fmt.Printf("v type: %T\n", v)
			for _, j := range v.([]interface{}) {
				for l, m := range j.(map[string]interface{}) {
					fmt.Printf("type l: %T\t type m: %T\n", l, m)
				}
			}
		}
	}

}
