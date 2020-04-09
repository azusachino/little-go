package fetcher

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("err occurred")
		return nil, errors.New("wrong status code: " + string(res.StatusCode))
	}
	all, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return all, nil
}
