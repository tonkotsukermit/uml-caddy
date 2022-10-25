package gen

import (
	"bytes"
	"io"
	"net/http"

	"github.com/andybalholm/brotli"
)

//GetPlantUMLPng retrieves the raw bytes of a png from a known plant uml generator url plantuml.com/plantuml/png/
func GetPlantUMLPng(puml string, url string) ([]byte , error) {

	buf := new(bytes.Buffer)
	
	brot := brotli.NewWriter(buf)

	_, err := brot.Write([]byte(puml))
	if err != nil {
		return nil, err
	}

	err = brot.Close()
	if err != nil {
		return nil, err
	}

	res, err := http.Get(url + "-1" + buf.String() )
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, err
}