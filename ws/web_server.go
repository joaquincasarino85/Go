package ws

import (
	"fmt"
	"net/http"
)

type WebServer struct {
	Url string
}

const Url_Rock string = "https://rock.com.ar"

func (obj WebServer) Connect() *http.Response {
	resp, err := http.Get(obj.Url)
	if err != nil {
		fmt.Println("error al leer pagina")
		return resp
	}

	return resp
}
