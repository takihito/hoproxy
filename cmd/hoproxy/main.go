package main

import (
	"flag"
	"fmt"

	"github.com/takihito/hoproxy"
)

func main() {
	configFile := flag.String("conf", "config.yml", "specify config file")
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()

	if showVersion {
		fmt.Printf("hoproxy %v \n", hoproxy.Version)
		return
	}

	cfg := hoproxy.NewConfig(*configFile)
	fmt.Printf("hoproxy c %v \n", cfg)
	hoproxy.Run(cfg)
}
