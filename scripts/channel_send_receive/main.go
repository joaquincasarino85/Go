package main

import (
	"fmt"
	"project/scrapper/db"
	"project/scrapper/lib"
	"project/scrapper/ws"

	"github.com/PuerkitoBio/goquery"

	_ "github.com/go-sql-driver/mysql"
)

func send(c chan<- [][]lib.Artist) {

	doc := lib.GetHtmlContent(ws.Url_Rock)
	// Get artist list ids
	fmt.Println("Getting artists")
	artistsList := [][]lib.Artist{}
	div := doc.Find(".abctop")
	div.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			artistsList = append(artistsList, lib.GetArtists("https://rock.com.ar"+href))
			fmt.Printf(".")
		}
	})

	c <- artistsList
}

func read(c <-chan [][]lib.Artist) {

	for _, v := range <-c {
		lib.ProcessArtists(v)
	}
}

func main() {

	db.ConfigureDatabase()

	c := make(chan [][]lib.Artist)
	go send(c)

	read(c)
	fmt.Println("\n...End")

}
