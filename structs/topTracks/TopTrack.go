package topTracks

import "lastfm-grabber/structs/basic"

type TopTrack struct {
	Attr      TopTrackAttr `json:"@attr"`
	PlayCount string       `json:"playcount"`
	Artist    basic.Artist `json:"artist"`
	Name      string       `json:"name"`
}
