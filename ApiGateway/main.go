package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const serverPort = ":3250"
const textGenUrl = "http://127.0.0.1:5001/api/generate"

func apiStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Health check - ALIVE"))
	return
}

func generateText(w http.ResponseWriter, r *http.Request) {

	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseString := string(responseData)
	log.Println("Request body: ", responseString)

	resp, err := http.Post(textGenUrl, "text/plain", strings.NewReader(responseString))
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	log.Println("resp: ", resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(sb))
	return
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/status", apiStatus)
	mux.HandleFunc("/api/generateText", generateText)

	s := &http.Server{
		Addr:           serverPort,
		Handler:        mux,
		ReadTimeout:    40 * time.Second,
		WriteTimeout:   40 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
