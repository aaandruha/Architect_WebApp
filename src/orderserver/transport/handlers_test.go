import "testing"

func TestList(t *testing.T) {
	w := httptest.NewRecorder()
	list(w, nil)
	response := w.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Status error")
	}

	jsonString, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	items := make([]OrderListItem, 10)
	if err = json.Unmarshal(jsonString, &items); err != n {
		t.Errorf("1111%v", err)
	}
}