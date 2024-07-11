package main

import (
	"cosmos-ibc-test/core"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
)

type Config struct {
	Addresses map[string]string `json:"addresses"`
}

func loadConfig(configFile string) (Config, error) {
	var config Config
	file, err := os.Open(configFile)
	if err != nil {
		return config, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return config, err
	}

	if err := json.Unmarshal(byteValue, &config); err != nil {
		return config, err
	}

	return config, nil
}

func main() {

	// Define a flag for the config file
	configFile := flag.String("config", "addresses.json", "Path to the configuration file")
	flag.Parse()

	// if the config file is not provided, assume that it's on the root of the project
	if *configFile == "" {
		*configFile = "./addresses.json"
	}

	// Load the configuration
	config, err := loadConfig(*configFile)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	tokenSvc := core.NewTokenService(core.GetChainEndpoints(), core.GetOriginalDenoms())
	print := core.NewPrintService()

	tokenInfos, totalAmounts := tokenSvc.GetTokenInfo(config.Addresses)

	// print this
	print.PrintTokenInfo(tokenInfos, totalAmounts)
}
