package main

import (
	// "encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	// "bytes"
	"strings"
	"time"
)

const serverPort = ":3250"
const textGenUrl = "http://127.0.0.1:5001/api/generate"

func apiStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("Health check - ALIVE")
}

func generateText(w http.ResponseWriter, r *http.Request) {

	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseString := string(responseData)
	log.Println("RS: ", responseString)

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

	// log.Println("2222222")

	// tr := &http.Transport{
	// 	MaxIdleConns:       10,
	// 	IdleConnTimeout:    40 * time.Second,
	// 	DisableCompression: true,
	// }
	// client := &http.Client{Transport: tr}
	// log.Println("3333333")

	// resp, _ := client.Post(textGenUrl, "text/plain", strings.NewReader(text.InputText))
	// log.Println("44444")
	// if err != nil {
	// 	log.Fatalf("An Error Occured %v", err)
	// 	return
	// }
	// log.Println("55555")
	// // defer resp.Body.Close()
	// log.Println("66666")
	// b, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println("Error: ", err)
	// }
	// log.Println("77777")
	// log.Println("Response: ", string(b))
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
