package transport

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func Router() http.Handler { //*mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	//s.HandleFunc("/hello-world", helloWorld).Methods(http.MethodGet)
	s.HandleFunc("/hello/{id}/", getKitty).Methods(http.MethodGet)
	return logMiddleware(r)
}

func helloWorld(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		t := time.Now()
		log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent":  r.UserAgent(),
			"speed":      t.Sub(start),
		}).Info("got a new request")

	})
}

type Kitty struct {
	Name string `json:"name"`
}

func getKitty(w http.ResponseWriter, _ *http.Request) {
	cat := Kitty{"Кот"}
	b, _ := json.Marshal(cat)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(b))
}
