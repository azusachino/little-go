package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	roomFs, err := os.OpenFile("room.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer roomFs.Close()
	http.HandleFunc("/room", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		roomFs.Write(body)
		roomFs.WriteString("\r\n")
	})

	mediaFs, err := os.OpenFile("media.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer mediaFs.Close()
	http.HandleFunc("/media", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		mediaFs.Write(body)
		mediaFs.WriteString("\r\n")
	})

	xmergFs, err := os.OpenFile("xmerg.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer xmergFs.Close()
	http.HandleFunc("/xmerg", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		xmergFs.Write(body)
		xmergFs.WriteString("\r\n")
	})

	log.Println("server is listening 8785")

	http.ListenAndServe(":8785", nil)
}
