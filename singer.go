package music

import (
	"fmt"
	"time"
)

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

func (s *Singer) Sing(song Song) error {
	for _, v := range song.Lyrics {
		fmt.Printf("[%s] %s\n", s.Name, v)
	}
	return nil
}

func (s *Singer) SingPart(song Song, start, end int, ch chan string) {
	for _, v := range song.Lyrics[start:end] {
		fmt.Printf("[%s] %s\n", s.Name, v)
		time.Sleep(10 * time.Millisecond)
		ch <- v
	}
	close(ch)
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
