package main

import (
	"context"
	"net/http"
	"orderserver/pkg/orderservice/config"
	"orderserver/pkg/orderservice/transport"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

const (
	appID              = "orderserver"
	configRelativePath = "./config.json"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("my.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
		defer file.Close()
	}
	c, err := config.ParseEnv()
	if err == nil {
		log.Info("No ENV args")
	}
	serverURL := c.ServeRESTAddress

	log.WithFields(log.Fields{"url": serverURL}).Info("starting the server")
	router := transport.Router()
	log.Fatal(http.ListenAndServe(serverURL, router))
	killSignalChan := getKillSignalChan()
	srv := startServer(serverURL)

	waitForKillSignal(killSignalChan)
	srv.Shutdown(context.Background())
}

func startServer(serverURL string) *http.Server {
	router := transport.Router()
	srv := &http.Server{Addr: serverURL, Handler: router}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	return srv
}

func getKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitForKillSignal(killSignalChan <-chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Info("got SIGINT")
	case syscall.SIGTERM:
		log.Info("get SEGTERM")
	}
}

/*	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello World!")
			})

			http.HandleFunc("/api/v1/orders", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "[{\n\t\"id\": \"1234-736363-sjdhjuy76234-9282\",\n\t\"menuItems\":[{\n\t\t\"id\": \"111111-jjdjjd-dkdkkd\",\n\t\t\"quantity\": 1 \n\t}]\n}]")
			})
	ice
			http.HandleFunc("/api/v1/order/1234-736363-sjdhjuy76234-9282", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "{\n\t\"id\": \"1234-736363-sjdhjuy76234-9282\",\n\t\"menuItems\":[{\n\t\t\"id\": \"111111-jjdjjd-dkdkkd\",\n\t\t\"quantity\": 1 \n\t}]\n\t\"orderedAtTimestamp\": 165423372625\n\t \"cost\": 999\n}")
			})*/
