package main

import "go.uber.org/zap"

func main() {
	sample()
}

func sample() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()

	sugar.Infow("failed to fetch URL",
		"url", "http://azusachino.cn",
		"attempt", 3)

	sugar.Infof("Failed again: %s", "abba")
}
