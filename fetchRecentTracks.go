package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lastfm-grabber/structs/recentTracks"
	"net/http"
	"os"
)

func getRecent(user string) recentTracks.ListenData {
	key := os.Getenv("API_KEY")
	response, err := http.Get(fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=%v&api_key=%v&format=json", user, key))
	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	body, readError := ioutil.ReadAll(response.Body)

	if readError != nil {
		panic(readError)
	}
	var data recentTracks.ListenData
	parseError := json.Unmarshal(body, &data)
	if parseError != nil {
		fmt.Println("error:", err)
	}
	return data
}
