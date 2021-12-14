package music

import "time"

type Album struct {
	Name        string
	ReleaseDate time.Time
}

type Singer struct {
	Person
	IsAvailable bool

	albums []*Album
}

func (s *Singer) CreateAlbum(name string) *Album {
	album := &Album{
		Name: name,
	}
	s.albums = append(s.albums, album)
	return album
}

func (s *Singer) Available() bool {
	return s.IsAvailable
}

func NewSinger(name string) *Singer {
	return &Singer{
		Person: Person{
			Name: name,
		},
		IsAvailable: true,
	}
}

type ProfessionalSinger interface {
	Available() bool
}
