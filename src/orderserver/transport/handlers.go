package transport

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	//s.HandleFunc("/hello-world", helloWorld).Methods(http.MethodGet)
	//s.HandleFunc("/hello/{id}/", getKitty).Methods(http.MethodGet)
	s.HandleFunc("/order", createOrder).Methods(http.MethodPost)
	return r //ogMiddleware(r)
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//TODO

	}
	defer r.Body.Close()
	var msg Order
	err = json.Unmarshal(b, &msg)
	if err != nil {
		//TODO
	}
	// check order
	fmt.Println(msg.Id)

	if msg.Id != "2424322" {
		http.Error(w, "my own error message", http.StatusForbidden)
		return
	}

	id := uuid.New()
	orderQuery := Order{Id: id.String(), menuItems: []Menu{{Id: "12345", Quantity: 3}, {Id: "12333345", Quantity: 2}}}

	b1, err := json.Marshal(orderQuery)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, string(b1)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
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

type Order struct {
	Id        string `json:"id"`
	menuItems []Menu `json:"menuItems"`
}

type Menu struct {
	Id       string `json:"id"`
	Quantity int    `json:"quantity"`
}

func getKitty(w http.ResponseWriter, _ *http.Request) {
	cat := Kitty{"Кот"}
	b, _ := json.Marshal(cat)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(b))
}

/*type Server struct {
	db *sql.DB
}

func (s *Server) createOrder(http.ResponseWriter, *http.Request) {

	db, err := s.db.Open("mysql", `root`)
	result, err := s.db.Exec()



}*/
