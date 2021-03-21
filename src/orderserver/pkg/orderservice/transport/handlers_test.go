package transport

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestList(t *testing.T) {
	/*	w := httptest.NewRecorder()
		//list(w, nil)
		response := w.Result()
		if response.StatusCode != http.StatusOK {
			t.Errorf("Status error")
		}

		jsonString, err := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			t.Fatal(err)
		}
		/*items := make([]Order, 10)
		if err = json.Unmarshal(jsonString, &items); err != nil {
			t.Errorf("%v", err)
		}*/

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}

	req := httptest.NewRequest("GET", "http://localhost:8000/api/v1/order/926c9a76-4464-11eb-bdf0-ee331b8c8f24", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	response := w.Result()
	jsonString, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		t.Errorf("Status error")
	}
	response.Body.Close()
	items := make([]Order, 1)
	if err := json.Unmarshal(jsonString, &items); err != nil {
		t.Errorf("%v", err)
	}

}
