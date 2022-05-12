package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/DazWilkin/go-buymeacoffee/types"

	"github.com/gorilla/mux"
)

const (
	purchaseID     string = "2621"
	subscriptionID string = "7979"
	supporterID    string = "245731"
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

// Content is a type that represents the data that is passed to the Golang Template
type Content struct {
	Handlers map[string]string
}

var (
	// Define the mapping of handlers for the Golang Template
	handlers = map[string]string{
		"/extras":             "/extras",
		"/extras/{id}":        fmt.Sprintf("/extras/%s", purchaseID),
		"/subscriptions":      "/subscriptions",
		"/subscriptions/{id}": fmt.Sprintf("/subscriptions/%s", subscriptionID),
		"/supporters":         "/supporters",
		"/supporters/{id}":    fmt.Sprintf("/supporters/%s", supporterID),
	}
	// Define the instance of Content used by handleRoot
	content = Content{
		Handlers: handlers,
	}
)

func handleRoot(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	t := template.Must(template.New("root.tmpl").ParseFiles("root.tmpl"))
	if err := t.ExecuteTemplate(w, "content", content); err != nil {
		log.Fatal(err)
	}
}
func handleAllExtras(w http.ResponseWriter, _ *http.Request) {
	w.Write(types.ExamplePurchases)
}
func handleOneExtras(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ID := vars["id"]

	if ID == purchaseID {
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

	if ID == subscriptionID {
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

	if ID == supporterID {
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

	r.HandleFunc("/", handleRoot)
	r.HandleFunc("/extras", handleAllExtras).Methods("GET")
	r.HandleFunc("/extras/{id}", handleOneExtras).Methods("GET")
	r.HandleFunc("/subscriptions", handleAllSubscriptions).Methods("GET")
	r.HandleFunc("/subscriptions/{id}", handleOneSubscriptions).Methods("GET")
	r.HandleFunc("/supporters", handleAllSupporters).Methods("GET")
	r.HandleFunc("/supporters/{id}", handleOneSupporters).Methods("GET")

	log.Printf("Starting server (%s)", *endpoint)
	log.Println(http.ListenAndServe(*endpoint, r))
}
