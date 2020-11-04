package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lastfm-grabber/structs/topTracks"
	"net/http"
	"os"
)

func getTop(user string) topTracks.ListenData {
	key := os.Getenv("API_KEY")
	response, err := http.Get(fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=user.gettoptracks&user=%v&api_key=%v&format=json", user, key))

	if err != nil {
		panic(err)
	}

	body, readError := ioutil.ReadAll(response.Body)

	if readError != nil {
		panic(readError)
	}

	var data topTracks.ListenData
	parseError := json.Unmarshal(body, &data)
	if parseError != nil {
		fmt.Println("error:", err)
	}
	return data
}
