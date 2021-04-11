package transport

import (
	"database/sql"
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
	s.HandleFunc("/orders", list).Methods(http.MethodGet)
	s.HandleFunc("/order", createOrderSql).Methods(http.MethodPost)
	s.HandleFunc("/order/{uuid:[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}}", getOrder).Methods(http.MethodGet)
	return r
}

func list(w http.ResponseWriter, r *http.Request) {
	ordersList := [3]Order{
		{Id: "926c9a76-4464-11eb-bdf0-ee331b8c8f24", MenuItems: []Menu{{Id: "444", Quantity: 3}, {Id: "12333345", Quantity: 2}}, OrderedAtTimestamp: 165423372625, Cost: 999},
		{Id: "926c9a76-4464-11eb-bdf0-ee333b8c8f24", MenuItems: []Menu{{Id: "12333345", Quantity: 2}, {Id: "12333345", Quantity: 1}}, OrderedAtTimestamp: 165423378625, Cost: 345},
		{Id: "926234a76-4464-112b-bdf0-ee331b8c8f24", MenuItems: []Menu{{Id: "12333345", Quantity: 1}, {Id: "12333345", Quantity: 3}}, OrderedAtTimestamp: 165423322625, Cost: 550},
	}
	b1, err := json.Marshal(ordersList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err := io.WriteString(w, string(b1)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuidOrder := vars["uuid"]

	orderQuery := Order{Id: uuidOrder, MenuItems: []Menu{{Id: "12345", Quantity: 3}, {Id: "12333345", Quantity: 2}}, OrderedAtTimestamp: 165423372625, Cost: 999}
	b1, err := json.Marshal(orderQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err := io.WriteString(w, string(b1)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	var request Order
	err = json.Unmarshal(b, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(request.MenuItems) == 0 {
		http.Error(w, "MenuItem should be greater than zero ", http.StatusBadRequest)
		return
	}
	//check quantity for menuitem
	var quantity int
	for _, s := range request.MenuItems {
		quantity += s.Quantity
	}
	if quantity == 0 {
		http.Error(w, "Quantity MenuItem should be greater than zero ", http.StatusBadRequest)
		return
	}

	id := uuid.New()
	orderQuery := Order{Id: id.String(), MenuItems: request.MenuItems, OrderedAtTimestamp: time.Now().UnixNano(), Cost: 1234}

	b1, err := json.Marshal(orderQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, string(b1)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
}

/*func helloWorld(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello world")
}*/

/*func logMiddleware(h http.Handler) http.Handler {
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
}*/

type Order struct {
	Id                 string `json:"id"`
	MenuItems          []Menu `json:"menuItems"`
	OrderedAtTimestamp int64  `json:"orderedAtTimestamp"`
	Cost               int    `json:"cost"`
}

type Menu struct {
	Id       string `json:"id"`
	Quantity int    `json:"quantity"`
	Cost     int    `json:"cost"`
}

/*func getKitty(w http.ResponseWriter, _ *http.Request) {
	cat := Kitty{"Кот"}
	b, _ := json.Marshal(cat)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(b))
}*/

type Server struct {
	db *sql.DB
}

func createOrderSql(http.ResponseWriter, *http.Request) {

	db, err := sql.Open("mysql", `root:root@/orderserver`)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	id := uuid.New()
	cost := 123.56
	q := "INSERT INTO `orderserver`.`order` (id, created, cost) VALUES (?, ?, ?)"
	fmt.Println(q)
	_, err = db.Exec(q, id.String(), time.Now(), cost)

	if err != nil {
		log.Fatal(err)
	}
}
