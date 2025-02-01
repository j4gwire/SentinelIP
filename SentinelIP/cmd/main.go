package main

import (
	"fmt"
	"log"
	"net/http"

	"SentinelIP/internal/proxy"
	"SentinelIP/internal/threat"
)

func main() {
	// Example to initialize threat and proxy detection modules
	threatDetector := threat.NewDetector()
	proxyDetector := proxy.NewDetector()

	// Example: Start an HTTP server for the service
	http.HandleFunc("/threat-detection", func(w http.ResponseWriter, r *http.Request) {
		// Example of using threat detection
		ip := r.RemoteAddr // Getting IP from the request
		isThreat, err := threatDetector.IsThreat(ip)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if isThreat {
			fmt.Fprintf(w, "IP %s is a threat!\n", ip)
		} else {
			fmt.Fprintf(w, "IP %s is safe.\n", ip)
		}
	})

	http.HandleFunc("/proxy-detection", func(w http.ResponseWriter, r *http.Request) {
		// Example of using proxy detection
		ip := r.RemoteAddr // Getting IP from the request
		isProxy, err := proxyDetector.IsProxy(ip)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if isProxy {
			fmt.Fprintf(w, "IP %s is using a proxy.\n", ip)
		} else {
			fmt.Fprintf(w, "IP %s is not using a proxy.\n", ip)
		}
	})

	// Start the HTTP server
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
