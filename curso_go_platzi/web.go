package main

import (
	"fmt"
	"io"
	"net/http"
)

type escritorWeb struct{}

func (escritorWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	respuesta, error := http.Get("http://google.com")
	if error != nil {
		fmt.Println(error)
	}
	e := escritorWeb{}
	io.Copy(e, respuesta.Body)
	fmt.Println(respuesta.Body)

}
