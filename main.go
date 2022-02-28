package main

import (
	"fmt"
	"project/scrapper/db"
	"project/scrapper/ws"

	"github.com/PuerkitoBio/goquery"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db.OpenConnection()
	db.CreateDatabase()

	doc := getHtmlContent(ws.Url_Rock)

	// LLeno la lista de ids de artistas
	fmt.Println("Getting artists")
	artistsList := [][]Artist{}
	div := doc.Find(".abctop")
	div.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			artistsList = append(artistsList, getArtists("https://rock.com.ar"+href))
			fmt.Printf(".")
		}
	})
	fmt.Println("OK!")

	processEntities(artistsList)
	fmt.Println("\n...End")
}
