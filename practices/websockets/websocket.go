package websockets

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

var count = flag.Int("count", 10, "client count")

const uri = "ws://localhost:8080/ws"

func Ws() {
	flag.Parse()
	log.SetFlags(0)

	// graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	// groutine manager
	stopCh := make(chan struct{})
	var wg sync.WaitGroup

	for i := 0; i < *count; i++ {
		wg.Add(1)
		go NewWebsocket(i, &wg, stopCh)
	}

	<-sigCh
	log.Println("received stop signal, ready to stop all sub-goroutine")

	close(stopCh)

	wg.Wait()
	log.Println("mission complete")
}

func NewWebsocket(index int, wg *sync.WaitGroup, stopCn <-chan struct{}) {
	defer wg.Done()

	c, _, err := websocket.DefaultDialer.Dial(uri, nil)
	if err != nil {
		log.Fatal("dail:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	// new goroutine to receive message
	go func() {
		defer close(done)
		for {
			_, raw, err := c.ReadMessage()
			if err != nil {
				log.Println("read: ", err)
				return
			}
			log.Printf("msg: %s", string(raw))
		}
	}()

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte("ping"))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-stopCn:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
