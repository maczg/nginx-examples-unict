package main

import (
	"fmt"
	mux "github.com/gorilla/mux"
	log "github.com/sirupsen/logrus" // use module logrus as log
	"net/http"
	"os"
)

var serverPort, hostname string

func init() {
	if serverPort = os.Getenv("SERVER_PORT"); serverPort == "" {
		// serverPort with env variable not set, use default port 80
		serverPort = "80"
	}
	// format server port with colon (:)
	serverPort = fmt.Sprintf(":%s", serverPort)

	//set hostname as global var
	hostname, _ = os.Hostname()

	// create new custom formatter and assign to default logger
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
}

// create home handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "%s's Homepage\n", hostname)
}

// create whoami handler
func whoamiHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = fmt.Fprintf(w, "I'm %s\n", hostname)
}

// create Greeting handler
func greetingHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bodyResp := fmt.Sprintf("Hi %s\n", q.Get("name"))
	_, _ = w.Write([]byte(bodyResp))
}

func main() {
	log.Infoln("..Starting server...")

	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/home", homeHandler).Methods("GET")
	router.HandleFunc("/whoami", whoamiHandler).Methods("GET")
	router.HandleFunc("/greeting", greetingHandler).Methods("POST")

	log.Errorln(http.ListenAndServe(serverPort, router))
}
