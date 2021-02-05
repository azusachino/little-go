package decorator

import (
	"fmt"
	"log"
	"net/http"
	_ "strings"
)

func WithServerHeader(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("--->WithServerHeader_()")
		w.Header().Set("Server", "HelloServer v0.0.1")
		h(w, r)
	}
}
func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Recieved Request %s from %s\n", r.URL.Path, r.RemoteAddr)
	_, _ = fmt.Fprintf(w, "Hello, World! "+r.URL.Path)
}
func main() {
	http.HandleFunc("/v1/hello_world", WithServerHeader(hello))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
