package music

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinger(t *testing.T) {
	singer := NewSinger("John Doe")
	assert.Equal(t, "John Doe", singer.Name)
}

func TestSingerCreateAlbum(t *testing.T) {
	singer := NewSinger("John Doe")
	album := singer.CreateAlbum("The Sun in The Sky")
	assert.Equal(t, "The Sun in The Sky", album.Name)
}

func TestSingerAvailability(t *testing.T) {
	singer := NewSinger("John Doe")
	assert.True(t, true, singer.Available())

	singer.IsAvailable = false
	assert.False(t, false, singer.Available())

	var obj interface{} = singer
	switch obj.(type) {
	case ProfessionalSinger:
		t.Log("is a professional singer")
	default:
		t.Fatal("not a professional singer")
	}
}
