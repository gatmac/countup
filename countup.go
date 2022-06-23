package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

const CONFIG_FN = ".countup"

func getConfig(filename string) map[string]time.Time {
	configMap := make(map[string]time.Time)
	osFile, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening %s: %v", filename, err)
	}
	csv, err := csv.NewReader(osFile).ReadAll()
	if err != nil {
		log.Fatalf("Error processing %s: %v", filename, err)
	}
	for _, row := range csv {
		if row[0][0] == '#' {
			continue
		}
		date, err := time.Parse("2006-01-02", row[0])
		if err != nil {
			log.Fatalf("Error processing date %s: %v", row[0], err)
		}
		configMap[row[1]] = date
	}
	return configMap
}

func getEnvironmentVariables() map[string]string {
	if len(os.Environ()) == 0 {
		log.Fatalln("No environment variables found.")
	}
	env := make(map[string]string)
	for _, e := range os.Environ() {
		for i, c := range e {
			if c == '=' {
				env[e[0:i]] = e[i+1:]
			}
		}
	}
	return env
}

func main() {
	now := time.Now()
	env := getEnvironmentVariables()
	cfgFile := env["HOME"] + "/" + CONFIG_FN
	//log.Printf("Using %s", cfgFile)
	cfg := getConfig(cfgFile)
	for k, v := range cfg {
		diff := math.Round(now.Sub(v).Hours() / 24)
		fmt.Printf("Day %.0f of %s\n", diff, k)
	}
}
