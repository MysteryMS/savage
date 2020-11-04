package recentTracks

type RecentTracks struct {
	Attr   RecentAttr `json:"@attr"`
	Tracks []Track    `json:"track"`
}
