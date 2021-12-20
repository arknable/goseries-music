package music

import "time"

type Song struct {
	Lyrics []string
}

type Album struct {
	Name        string
	ReleaseDate time.Time
}
