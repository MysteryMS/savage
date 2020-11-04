package recentTracks

import "lastfm-grabber/structs/basic"

type Track struct {
	Artist basic.Artist `json:"artist"`
	Attr   TrackAttr    `json:"@attr"`
	Album  basic.Album  `json:"album"`
	Name   string       `json:"name"`
}
