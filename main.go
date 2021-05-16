package main

import (
	"fmt"
	"log"
	"time"
)

const (
	dateFormat = "2006-01-02"
)

type Holiday struct {
	start string
	end   string
}

var Vacation = []Holiday{
	{start: "2020-01-01", end: "2020-01-04"},
	{start: "2020-03-01", end: "2020-03-31"},
	// {start: "2021-05-06", end: "2020-05-06"},
}

func ComputeDays(v Holiday) float64 {
	start, err := time.Parse(dateFormat, v.start)
	if err != nil {
		log.Fatal(err)
	}

	end, err := time.Parse(dateFormat, v.end)
	if err != nil {
		log.Fatal(err)
	}
	duration, err := time.ParseDuration(end.Sub(start).String())
	if err != nil {
		log.Fatal(err)
	}
	convertToDays := duration.Hours() / 24.0
	if convertToDays < 0 {
		log.Fatal("your vacation can't have a negative length")
	}

	return convertToDays
}

func main() {
	// fmt.Println(Vacation)
	var totalDays float64
	for _, v := range Vacation {
		daysAway := ComputeDays(v)
		totalDays += daysAway
	}

	fmt.Printf("You have spent %v days out of the UK", totalDays)
}
