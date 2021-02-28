package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var paths = map[string]string{
	"/go-path":   "http://yandex.ru",
	"/go-gopher": "https://gopl.io",
}

func main() {
	http.HandleFunc("/", route)
	http.ListenAndServe(":8090", nil)
}

func route(w http.ResponseWriter, r *http.Request) {

	if len(os.Args) > 1 && os.Args[1] == "-r" {
		var err error
		paths, err = parseShortenURL(os.Args[2])
		if err != nil {
			return
		}
	}

	if len(paths[r.URL.Path]) > 0 {
		// redirect
		http.Redirect(w, r, paths[r.URL.Path], http.StatusSeeOther)
		return
	}
	// fallback message
	fmt.Fprintf(w, "Искомый URL %s не найден\n", r.URL.Path)
}

func parseShortenURL(fileName string) (map[string]string, error) {

	var jsonResult map[string]string
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &jsonResult)

	return jsonResult, nil
}
