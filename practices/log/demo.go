package main

import (
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger
var file os.File

func Init() {
	encoder := getEncoder()
	defaultWriter := getLogWriter("E:\\Projects\\project-github\\little-go\\practices\\log\\logs")
	core := zapcore.NewCore(
		encoder, defaultWriter, zap.LevelEnablerFunc(
			func(l zapcore.Level) bool {
				return l >= zapcore.DebugLevel
			}),
	)
	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func main() {
    Init()
	defer sugarLogger.Sync()
	defer file.Close()
	url := "https://azusachino.cn"
	write(&url)
}

func write(url *string) {
	done := make(chan bool)

	go func() {
		time.Sleep(time.Minute)
		done <- true
	}()

	for {
		select {
		case <-done:
			return
		default:
			sugarLogger.Infof("failed to fetch URL %s", *url)
		}
	}
}

func getEncoder() zapcore.Encoder {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(cfg)
}

func getLogWriter(filePath string) zapcore.WriteSyncer {
	writer := getWriter(filePath)
	return zapcore.AddSync(writer)
}

func getWriter(filePath string) io.Writer {
	today := time.Now().Format("2006-01-02")
	file, err := os.Create(filePath + "\\" + today + ".log")
	if err != nil {
		panic(err)
	}
	return file
}

func sample() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	url := "https://azusachino.cn"

	sugar := logger.Sugar()

	sugar.Infow("failed to fetch URL",
		"url", "http://azusachino.cn",
		"attempt", 3)

	sugar.Infof("Failed again: %s", "abba")

	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
