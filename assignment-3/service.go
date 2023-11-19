package main

import (
	"log"
	"math/rand"
	"os"
	"time"
)

func randomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n) + 1
}

func getStatus(value, min, max string) string {
	switch {
	case value < min:
		return "aman"
	case value >= min && value <= max:
		return "siaga"
	case value > max:
		return "bahaya"
	default:
		return "status tidak diketahui"
	}
}

func logToFile(filename string, message string) {
	logFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.Println(message)
}
