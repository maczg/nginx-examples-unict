package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

func main() {
	var port string
	if port = os.Getenv("WH_PORT"); port == "" {
		port = "80"
	}

	hostname, _ := os.Hostname()

	log.Infoln("Hostname: ", hostname, " server starting on port ", port)

	router := mux.NewRouter()
	router.HandleFunc("/whoami", whoamiHandler).Methods("GET")
	router.HandleFunc("/invoke", invokeHandler).Methods("GET")

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

// create whoami handler
func whoamiHandler(w http.ResponseWriter, r *http.Request) {

	// get hostname
	hostname, err := os.Hostname()
	if err != nil {
		abortWithMessage(w, http.StatusInternalServerError, "Error getting hostname")
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = fmt.Fprintf(w, "Hostname is: %s\n", hostname)
}

type invokeRequest struct {
	Endpoint string `json:"endpoint"`
	Path     string `json:"path"`
}

func invokeHandler(w http.ResponseWriter, r *http.Request) {
	var request invokeRequest
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		abortWithMessage(w, http.StatusInternalServerError, "Error reading request body")
	}

	err = json.Unmarshal(raw, &request)
	if err != nil {
		abortWithMessage(w, http.StatusInternalServerError, "Error unmarshalling request body")
	}
	url := fmt.Sprintf("http://%s/%s", request.Endpoint, request.Path)
	log.Infoln("Invoking ", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		abortWithMessage(w, http.StatusInternalServerError, "Error creating request")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		abortWithMessage(w, http.StatusInternalServerError, "Error invoking endpoint")
	}
	log.Infoln("Response from ", url, ": ", resp.Status)
	rawResp, _ := io.ReadAll(resp.Body)
	_, _ = w.Write(rawResp)
}

func abortWithMessage(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	_, _ = fmt.Fprintf(w, "%s\n", message)
}
