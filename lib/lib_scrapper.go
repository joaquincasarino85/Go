package lib

import (
	"fmt"
	"log"
	"project/scrapper/db"
	"project/scrapper/ws"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Song struct {
	title  string
	lyrics string
}

type Artist struct {
	id   string
	name string
}

func getWebServer(url string) ws.WebServer {
	return ws.WebServer{
		url,
	}
}

func GetHtmlContent(url string) *goquery.Document {
	ws := getWebServer(url)
	resp := ws.Connect()
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func GetSongLyrics(artistId string, songId string) Song {

	songLyricsHtml := GetHtmlContent(ws.Url_Rock + "/artistas/" + artistId + "/letras/" + songId)

	div := songLyricsHtml.Find("div.post-content-text")
	title := div.Find("h3").Text()

	return Song{
		title:  title,
		lyrics: div.Find("div").Eq(1).Text(),
	}
}

func GetSongs(artistId string) []string {

	songsHtml := GetHtmlContent(ws.Url_Rock + "/artistas/" + artistId + "/letras")
	songsList := []string{}
	songsHtml.Find("ul.canciones").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, s1 *goquery.Selection) {
			href, exists := s1.Attr("href")
			if exists {
				split := strings.Split(href, "/")
				songsList = append(songsList, split[4])
			}
		})
	})
	return songsList
}

func GetArtists(url string) []Artist {

	artistsHtml := GetHtmlContent(url)
	artistsList := []Artist{}
	artistsHtml.Find("ul.canciones").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(j int, s1 *goquery.Selection) {
			href, exists := s1.Attr("href")
			if exists {
				name, _ := s1.Html()
				artistsList = append(artistsList, Artist{
					id:   strings.Split(href, "/")[2],
					name: name,
				})
			}
		})
	})

	return artistsList
}

func ProcessArtists(artists []Artist) {

	for _, z := range artists {
		fmt.Println("Inserting artist: ", z.name)
		artistDbId, err := db.InsertArtist(z.name)
		if err != nil {
			log.Fatal(err)
		}

		songsList := GetSongs(z.id)
		if len(songsList) > 0 {
			fmt.Println("Inserting songs ")
		}
		for _, s := range songsList {
			songObj := GetSongLyrics(z.id, s)
			_, err := db.InsertSong(artistDbId, songObj.title, songObj.lyrics)
			if err != nil {
				log.Fatal(err)
			}
		}
		if len(songsList) > 0 {
			fmt.Println("OK!")
		}

	}

}
