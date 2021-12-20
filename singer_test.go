package music

import (
	"fmt"
	"sync"
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

func TestSingerSingASong(t *testing.T) {
	var (
		bambang = NewSinger("Bambang")
		cyntia  = NewSinger("Cyntia")

		bambangChan = make(chan string)
		cyntiaChan  = make(chan string)

		checkSingFunc = func(singer *Singer, ch chan string, wg *sync.WaitGroup) {
			defer wg.Done()

			for {
				v, ok := <-ch
				if !ok {
					fmt.Println(singer.Name, " is done")
					break
				}
				fmt.Printf("%s sang by %s\n", v, singer.Name)
			}
		}

		song = Song{
			Lyrics: []string{
				"Balonku", "ada", "lima",
				"Rupa", "rupa", "warnanya",
				"Hijau", "kuning", "kelabu",
				"Merah", "muda", "dan", "biru",
				"Meletus", "balon", "hijau,", "dor!",
				"Hatiku", "sangat", "kacau",
				"Balonku", "tinggal", "empat",
				"Kupegang", "erat-erat",
			},
		}
	)

	go bambang.SingPart(song, 0, 13, bambangChan)
	go cyntia.SingPart(song, 13, 25, cyntiaChan)

	var wg sync.WaitGroup
	wg.Add(2)
	go checkSingFunc(bambang, bambangChan, &wg)
	go checkSingFunc(cyntia, cyntiaChan, &wg)

	wg.Wait()
}
