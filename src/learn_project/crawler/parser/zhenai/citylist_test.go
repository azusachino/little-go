package zhenai

import (
	"learn_project/crawler/fetcher"
	"testing"
)

func test_a(t *testing.T) {
	contents, err := fetcher.Fetch("")
	if err != nil {
		panic(err)
	}
	ParseCityList(contents)
}
