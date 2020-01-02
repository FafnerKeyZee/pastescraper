package main

import (
	"context"
	"flag"
	"github.com/pimmytrousers/pastescraper/parse"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "./exampleconf.yml", "config file for the service")
}

func main() {
	flag.Parse()

	c := &config{}
	err := c.getConf(configPath)
	if err != nil {
		log.Fatalf("failed to acquire config: %s", err)
	}

	parser, err := parse.New(c.Parsers)
	if err != nil {
		log.Fatalf("failed to initialize parser: %s", err)
	}

	scraper, err := New(c, parser)
	if err != nil {
		log.Fatalf("failed to initialize scraper: %s", err)
	}

	err = scraper.start(context.Background(), time.Second * 10)
	if err != nil {
		log.Fatalf("failed to scrape: %s", err)
	}
}