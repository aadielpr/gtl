package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/aadielpr/gtl/country"
)

func main() {
	fmt.Println("Initial commit.")

	f, err := os.Open("countries.json")
	if err != nil {
		fmt.Println("Error when opening file: ", err)
		return
	}

	defer f.Close()

	byte, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("Error when reading file: ", err)
		return
	}

	var countries []country.Country

	err = json.Unmarshal(byte, &countries)
	if err != nil {
		fmt.Println("Error when Unmarshaling: ", err)
		return
	}

	jsonData, err := json.MarshalIndent(countries, "", " ")

	fmt.Println(string(jsonData))
}
