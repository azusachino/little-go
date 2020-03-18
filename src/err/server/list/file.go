package file

import (
	"io/ioutil"
	"net/http"
	"os"
)

func FileListing(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path[len("/list"):]
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
}
