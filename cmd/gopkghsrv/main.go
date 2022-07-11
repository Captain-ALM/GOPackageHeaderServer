package main

import "log"

var (
	buildVersion = "develop"
	buildDate    = ""
)

func main() {
	log.Printf("[Main] Starting up GO Package Header Server #%s (%s)\n", buildVersion, buildDate)
}
