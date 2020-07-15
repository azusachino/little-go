package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func start(p string, port string) {
	http.HandleFunc(p, func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path[len(p):]
		file, err := os.Open(path)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()
		all, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		writer.Write(all)
	})
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type UserError interface {
	error
	Message() string
}

type userError string

func (e userError) Error() {

}

func main() {
	start("/list", "9090")
}
