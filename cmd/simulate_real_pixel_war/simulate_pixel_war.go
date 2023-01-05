package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the CSV file
	// Get the dataset at https://www.reddit.com/r/place/comments/txvk2d/rplace_datasets_april_fools_2022/
	csvFile, err := os.Open("./cmd/simulate_real_pixel_war/place_history_sample.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer csvFile.Close()

	// Parse the CSV file
	reader := csv.NewReader(csvFile)

	// Skip the first line (column headers)
	reader.Read()

	// Send a new_place request to the server
	resp, err := http.Post(
		"http://localhost:8080/new_place",
		"application/json",
		strings.NewReader(fmt.Sprintf(`{"height":%d,"width":%d}`,
			2000,
			2000,
		)))

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// Unmarshal the response into a NewPlaceResponse struct
	var newPlaceResponse agent.NewPlaceResponse
	err = json.Unmarshal(body, &newPlaceResponse)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return
	}

	// The place ID is now stored in newPlaceResponse.PlaceID
	fmt.Printf("Place ID: %v\n", newPlaceResponse.PlaceID)
	placeID := newPlaceResponse.PlaceID

	if err != nil {
		fmt.Println(err)
		return
	}

	// Read each record
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}

		// Get the values from the record
		userID := record[1]
		pixelColor := record[2]
		coordinateStr := record[3]

		// Split the coordinate into x and y values
		coordinates := strings.Split(coordinateStr, ",")
		x, err := strconv.Atoi(coordinates[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		y, err := strconv.Atoi(coordinates[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		// Send a paint_pixel request to the server
		_, err = http.Post(
			"http://localhost:8080/paint_pixel",
			"application/json",
			strings.NewReader(fmt.Sprintf(`{"user-id":"%s","place-id":"%s","color":"%s","x":%d,"y":%d}`,
				userID,
				placeID,
				pixelColor,
				x,
				y,
			)))

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
