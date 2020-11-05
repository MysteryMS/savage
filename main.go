package main

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	if len(os.Args) == 1 {
		err := errors.New("missing \"user\" parameter\n")
		fmt.Print(err)
		return
	}
	user := "MysteryMS"
	fmt.Printf("Looking up for %v\n", user)

	data := getRecent(user)
	fmt.Printf("Found %v\n", color.CyanString(data.RecentTracks.Attr.User))

	firstTrack := data.RecentTracks.Tracks[0]
	if data.RecentTracks.Tracks[0].Attr.NowPlaying == "" {
		fmt.Printf(":: Last listened track 🡒 %v (by %v on %v)\n",
			color.BlueString(firstTrack.Name),
			color.MagentaString(firstTrack.Artist.Name),
			color.MagentaString(firstTrack.Album.Name),
		)
	} else {
		fmt.Printf("%v 🡒 %v by %v on %v\n",
			color.GreenString(":: Now Playing"),
			color.YellowString(firstTrack.Name),
			color.MagentaString(firstTrack.Artist.Name),
			color.MagentaString(firstTrack.Album.Name),
		)
	}

	color.Blue(":: Top 5 tracks ::")
	topTrackData := getTop(user)
	for pos := 0; pos < 5; pos++ {
		fmt.Printf("#%v - %v by %v (played %v times)\n",
			pos+1,
			color.BlueString(topTrackData.TopTracks.Tracks[pos].Name),
			color.MagentaString(topTrackData.TopTracks.Tracks[pos].Artist.Name),
			color.RedString(topTrackData.TopTracks.Tracks[pos].PlayCount),
		)
	}
}
