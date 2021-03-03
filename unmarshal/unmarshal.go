package main

import (
	"encoding/json"
	"fmt"
)

type pathsJSON struct {
	//Paths []string `json:"paths"`
	Paths path `json:"paths"`
}

//var paths map[string]interface{}

type path struct {
	Short string
}

func main() {
	jsonString := `{"paths":["/go-path","11111", "22222"]}` //`{"paths":["/go-path":"http://yandex.ru","/go-gopher":"https://gopl.io"]}`

	paths := pathsJSON{}

	err := json.Unmarshal([]byte(jsonString), &paths)

	if err != nil {
		return //errors.New("CustomFloat64: UnmarshalJSON: " + err.Error())
	}

	fmt.Println(paths.Paths)

	/*	for d, t := range paths.paths {
		fmt.Println(d, "-", t)
	}*/

}

func (u *path) UnmarshalJSON(data []byte) error {
	fmt.Println(data[0])
	err := json.Unmarshal(data, &u.Short)
	if err != nil {
		return err //errors.New("CustomFloat64: UnmarshalJSON: " + err.Error())
	}

	/*if data[0] == QUOTES_BYTE {
		err := json.Unmarshal(data[1:len(data)-1], &cf.Float64)
		if err != nil {
			return errors.New("CustomFloat64: UnmarshalJSON: " + err.Error())
		}
	} else {
		err := json.Unmarshal(data, &cf.Float64)
		if err != nil {
			return errors.New("CustomFloat64: UnmarshalJSON: " + err.Error())
		}
	}*/
	return nil
}
