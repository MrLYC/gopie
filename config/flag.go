package config

import (
	"flag"
	"os"
)

// Parse config
func Parse() {
	var help bool
	flag.BoolVar(&help, "h", false, "help message")
	flag.StringVar(&Address, "l", "0.0.0.0:8000", "listening address")
	flag.BoolVar(&Debug, "debug", false, "debug mode")
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(1)
	}
}
