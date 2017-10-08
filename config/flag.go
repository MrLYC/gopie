package config

import (
	"flag"
	"os"
)

// Parse config
func Parse() {
	var help bool
	flag.BoolVar(&help, "h", false, "help message")
	flag.StringVar(&Address, "l", Address, "listening address")
	flag.StringVar(&WorkDir, "w", WorkDir, "work directory")
	flag.BoolVar(&Debug, "debug", Debug, "debug mode")
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(1)
	}
}
