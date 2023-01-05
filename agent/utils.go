package agent

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

func ContainsHobby(hobbies []string, hobby string) bool {
	for _, v := range hobbies {
		if v == hobby {
			return true
		}
	}
	return false
}

func exists(s []*AgentWorker, id string) int {
	for k, v := range s {
		if v.id == id {
			return k
		}
	}
	return -1
}

func GetManagerIndex(s []*AgentManager, id string) int {
	for k, v := range s {
		if v.id == id {
			return k
		}
	}
	return -1
}

func remove[T comparable](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func MakeRandomSliceOfHobbies(hobbies []string) (result []string) {
	result = make([]string, 0)
	for i := 0; i < 1; i++ {
		k := rand.Intn(len(hobbies))
		result = append(result, hobbies[k])
	}
	return
}

func CreateNewPlace(url string, height, width int) (placeID string) {
	url = url + "/new_place"

	// Send a new_place request to the server
	resp, err := http.Post(
		url,
		"application/json",
		strings.NewReader(fmt.Sprintf(`{"height":%d,"width":%d}`,
			height,
			width,
		)))

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// Unmarshal the response into a NewPlaceResponse struct
	var newPlaceResponse NewPlaceResponse
	err = json.Unmarshal(body, &newPlaceResponse)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return
	}

	// The place ID is now stored in newPlaceResponse.PlaceID
	placeID = newPlaceResponse.PlaceID

	return placeID
}
