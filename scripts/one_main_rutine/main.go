package main

import (
	"fmt"
	"project/scrapper/db"
	"project/scrapper/lib"
	"project/scrapper/ws"

	"github.com/PuerkitoBio/goquery"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db.ConfigureDatabase()

	doc := lib.GetHtmlContent(ws.Url_Rock)

	// Get artist list ids  la lista
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
	fmt.Println("OK!")

	for _, v := range artistsList {
		lib.ProcessArtists(v)
	}

	fmt.Println("\n...End")

}
