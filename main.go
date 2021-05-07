package main

import (
	"context"
	"flag"
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var (
	logger = log.New()
	port   *string
	cache  *bool
	rdb    *redis.Client
	ctx    context.Context
)

func init() {
	// initialize logging
	logLevel := flag.Int("logLevel", 0, "defines the log level. 0=production builds. 1=dev builds.")
	port = flag.String("port", "8080", "port to bind http server to.")
	cache = flag.Bool("cache", false, "enable cache for coingecko")
	flag.Parse()
	logger.Out = os.Stdout
	switch *logLevel {
	case 0:
		logger.SetLevel(log.InfoLevel)
	default:
		logger.SetLevel(log.DebugLevel)
	}
}

func main() {
	var wg sync.WaitGroup

	if *cache {
		rdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		ctx = context.Background()
	}

	wg.Add(1)
	_ = NewManager(*port, rdb, ctx)

	// wait forever
	wg.Wait()
}
