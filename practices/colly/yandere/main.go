package main

import (
	"fmt"
)

const BASE_URL = "https://yande.re"

// 下载模式
type FetchType uint8

// 热门模式
type PopularType uint8

const (
	// 按热门下载
	Popular FetchType = iota
	// 按标签下载
	Tag
	// 按Id下载
	ShowId
)

const (
	Daily PopularType = iota
	Weekly
	Monthly
)

type YandreFetcher struct {
}

func DownloadByShowId(showId string) {
	// TODO
}

func main() {
	fmt.Println("hello yandere")
}
