package main

import (
	"fmt"
	"project/scrapper/db"
	"project/scrapper/lib"
	"project/scrapper/ws"
	"runtime"
	"sync"

	"github.com/PuerkitoBio/goquery"

	_ "github.com/go-sql-driver/mysql"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {

	db.ConfigureDatabase()

	doc := lib.GetHtmlContent(ws.Url_Rock)

	// LLeno la lista de ids de artistas
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

	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Go Routines will be:", len(artistsList))

	wg.Add(len(artistsList))
	for i := 0; i < len(artistsList); i++ {
		value := artistsList[i]
		go func() {
			mu.Lock()
			lib.ProcessArtists(value)
			runtime.Gosched()
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("\n...End")

}
