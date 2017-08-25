package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/wangdisdu/esbeat/beater"
)

var Version = "1.0.0"

func main() {
	err := beat.Run("esbeat", Version, beater.New)
	if err != nil {
		os.Exit(1)
	}
}
