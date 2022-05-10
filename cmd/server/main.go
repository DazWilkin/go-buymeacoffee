package main

import (
	"flag"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/DazWilkin/go-buymeacoffee/types"

	"github.com/gorilla/mux"
)

const (
	homePage string = `
<!DOCTYPE html>
<html lang="en-US">
<head>
  <title>Buy Me A Coffee -- Test API Server</title>
</head>
<body>
  <h1>Buy Me A Coffee -- Test API Server</h1>
  <hr/>
  <ul>
    <li>/extras</li>
	<li>/extras/{id}</li>
	<li>/subscriptions</li>
	<li>/subscriptions/{id}</li>
	<li>/supporters</li>
	<li>/supporters/{id}</li>
  </ul>
</body>
</html>
`
)

var (
	// BuildTime is the time that this binary was built represented as a UNIX epoch
	BuildTime string
	// GitCommit is the git commit value and is expected to be set during build
	GitCommit string
	// GoVersion is the Golang runtime version
	GoVersion = runtime.Version()
	// OSVersion is the OS version (uname --kernel-release) and is expected to be set during build
	OSVersion string
	// StartTime is the start time of the exporter represented as a UNIX epoch
	StartTime = time.Now().Unix()
)
var (
	endpoint = flag.String("endpoint", "0.0.0.0:8080", "Endpoint of server")
)

func home(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(homePage))
}
func handleAllExtras(w http.ResponseWriter, _ *http.Request) {
	w.Write(types.ExamplePurchases)
}
func handleOneExtras(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ID := vars["id"]

	if ID == "30" {
		w.Write(types.ExamplePurchase)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}
func handleAllSubscriptions(w http.ResponseWriter, _ *http.Request) {
	w.Write(types.ExampleSubscriptions)
}
func handleOneSubscriptions(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ID := vars["id"]

	if ID == "7979" {
		w.Write(types.ExampleSubscription)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}
func handleAllSupporters(w http.ResponseWriter, _ *http.Request) {
	w.Write(types.ExampleSupporters)
}
func handleOneSupporters(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ID := vars["id"]

	if ID == "245731" {
		w.Write(types.ExampleSupporter)
		return
	}

	w.WriteHeader(http.StatusNotFound)

}

func main() {
	flag.Parse()

	if *endpoint == "" {
		log.Fatal("Flag `--endpoint` is required")
	}

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/extras", handleAllExtras).Methods("GET")
	r.HandleFunc("/extras/{id}", handleOneExtras).Methods("GET")
	r.HandleFunc("/subscriptions", handleAllSubscriptions).Methods("GET")
	r.HandleFunc("/subscriptions/{id}", handleOneSubscriptions).Methods("GET")
	r.HandleFunc("/supporters", handleAllSupporters).Methods("GET")
	r.HandleFunc("/supporters/{id}", handleOneSupporters).Methods("GET")

	log.Printf("Starting server (%s)", *endpoint)
	log.Println(http.ListenAndServe(*endpoint, r))
}
