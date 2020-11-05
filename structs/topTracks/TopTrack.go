package topTracks

type TopTrack struct {
	Attr      TopTrackAttr   `json:"@attr"`
	PlayCount string         `json:"playcount"`
	Artist    TopTrackArtist `json:"artist"`
	Name      string         `json:"name"`
}
