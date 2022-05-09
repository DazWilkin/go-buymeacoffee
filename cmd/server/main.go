package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/DazWilkin/go-buymeacoffee/types"
)

var (
	endpoint = flag.String("endpoint", "0.0.0.0:8080", "Endpoint of server")
)

// ExtractID is function that parses candidate URL paths (/{method}/[{ExtractID}])
// It attempts to extract the method and an optional ExtractID
// An ExtractID of zero represents a valid path *without* an ExtractID i.e. /{method}/
func ExtractID(s string) (uint, error) {
	if s == "" {
		return 0, nil
	}

	ID, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(ID), nil
}
func handler(w http.ResponseWriter, req *http.Request) {
	// All API methods require GET
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Valid Paths are /METHOD or /METHOD/ID
	// Where METHOD is [purchases,subscriptions,supporters]
	// And ID parses as an unsigned integer
	path := req.URL.Path
	log.Printf("method: %s", path)

	// Examine the path by splitting it on "/"
	parts := strings.Split(path, "/")
	log.Printf("%d: %+v", len(parts), parts)

	if len(parts) < 2 || len(parts) > 4 {
		log.Println("Incorrect number of path components")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// If present, the method will always be part[1]
	method := parts[1]

	// If present and not empty, the ID will always be part[2]
	var ID uint
	var err error
	if len(parts) == 3 {
		if ID, err = ExtractID(parts[2]); err != nil {
			log.Printf("Unable to parse ID from part [%s]", parts[2])
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// Response are always JSON
	w.Header().Set("Content-Type", "application/json")

	switch method {
	case "extras":
		// If there's an ID, return /extras/{ID}
		if ID != 0 {
			// If ID matches ExamplePurchase ID==30
			if ID == 30 {
				w.Write(types.ExamplePurchase)
				return
			}

			w.WriteHeader(http.StatusNotFound)
			return
		}

		// Otherwise, return ExamplePurchases
		w.Write(types.ExamplePurchases)
		return
	case "subscriptions":
		// If there's an ID, return /subscriptions/{ID}
		if ID != 0 {
			// If ID matches ExampleSubscription ID==7979
			if ID == 7979 {
				w.Write(types.ExampleSubscription)
				return
			}

			w.WriteHeader(http.StatusNotFound)
			return
		}

		// Otherwise, return ExampleSubscriptions
		w.Write(types.ExampleSubscriptions)
		return
	case "supporters":
		// If there's an ID, return /subscriptions/{ID}
		if ID != 0 {
			// If ID matches ExampleSupporter ID==245731
			if ID == 245731 {
				w.Write(types.ExampleSupporter)
				return
			}
		}

		// Otherwise, return ExampleSupporters
		w.Write(types.ExampleSupporters)
		return
	default:
		log.Println("Unrecognized API method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	flag.Parse()

	if *endpoint == "" {
		log.Fatal("Flag `--endpoint` is required")
	}

	http.HandleFunc("/", handler)

	log.Printf("Starting server [%s]", *endpoint)
	log.Fatal(http.ListenAndServe(*endpoint, nil))
}
