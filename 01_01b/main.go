package main

import (
	"flag"
	"log"
	"time"
)

var expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {
	pt, err := time.Parse(expectedFormat, target)
	if err != nil || time.Now().After(pt) {
		log.Fatalln(err)
	}

	return pt
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	untilTarget := time.Until(target)

	return untilTarget.Hours() / 24.0
}

func main() {
	bday := flag.String("bday", "", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*bday)
	log.Printf("You have %d sleeps until your birthday. Hurray!",
		int(calcSleeps(target)))
}
