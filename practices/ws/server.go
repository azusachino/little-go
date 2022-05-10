package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrade = websocket.Upgrader{} // use default options

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection gradation:", err)
		return
	}
	defer conn.Close()

	// The event loop
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Received: %s", message)
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
	}
}

func home(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "Index Page")
}

func RunServer() {
	http.HandleFunc("/socket", socketHandler)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
