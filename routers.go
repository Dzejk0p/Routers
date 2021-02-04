package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"Routers/ping"

	"github.com/julienschmidt/httprouter"
)

func pinger(adresses []string) []string {
	result := make([]string, len(adresses))
	for i, address := range adresses {
		start := time.Now()
		reached := ping.Ping(address, 1)
		elapsed := time.Since(start).Milliseconds()
		row := address + "\t"
		if reached {
			row += fmt.Sprint(elapsed) + "ms"
		} else {
			row += "Unrechable"
		}
		result[i] = row
	}
	return result
}

func getHostname() (name string) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return
}

func getPings(req *http.Request) []string {
	var par string
	for key := range req.URL.Query() {
		par = key
	}

	adresses := []string{"r.pl", "google.com", "pornhub.com", "zieloneimperium.pl"}
	if par != "" {
		adresses = strings.Split(par, ",")
	}

	return pinger(adresses)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	fmt.Fprintf(w, "hostname: %s\n", getHostname())

	for _, address := range getPings(req) {
		fmt.Fprintf(w, "%s\n", address)
	}

	for _, env := range os.Environ() {
		fmt.Fprintln(w, env)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", index)

	log.Fatal(http.ListenAndServe(":8090", router))
}
