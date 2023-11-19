package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	ticker := time.NewTicker(15 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				water := randomNumber(100)
				wind := randomNumber(100)
				waterStatus := getStatus(strconv.Itoa(water), "5", "8")
				windStatus := getStatus(strconv.Itoa(wind), "6", "15")
				logMessage := fmt.Sprintf("\nWater: %d, Status: %s\nWind: %d, Status: %s", water, waterStatus, wind, windStatus)
				logToFile("log.txt", logMessage)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	// Simulate a delay before stopping the ticker
	time.Sleep(60 * time.Second)
	close(quit)
}
