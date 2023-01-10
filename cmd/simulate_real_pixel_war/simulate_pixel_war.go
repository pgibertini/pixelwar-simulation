package main

import (
	"encoding/csv"
	"fmt"
	agt "gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	// PARAMETERS
	url := "http://localhost:5555"

	// Open the CSV file
	// Get the full dataset at https://www.reddit.com/r/place/comments/txvk2d/rplace_datasets_april_fools_2022/
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

	// create a new place
	placeID := agt.CreateNewPlace(url, 2000, 2000, 0)

	fmt.Println("Basic front-end: ", url+"/canvas?placeID="+placeID)

	// Read each record
	i := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		if (i % 10000) == 0 {
			log.Printf("Pixel #%d\n", i)
		}
		i++

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
			url+"/paint_pixel",
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
