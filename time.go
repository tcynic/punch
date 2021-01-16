package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	// Retrieve local Unix Time
	localTime := time.Now().Local()

	// Make the time pretty
	prettyTime := localTime.Format("2006/01/02 15:04:05")

	// Create timelog file
	log, err := os.OpenFile("timelog", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print Date to timelog file
	_, err = fmt.Fprintln(log, prettyTime)
	if err != nil {
		fmt.Println(err)
		log.Close()
		return
	}

	// Close timelog file if ther is an error
	err = log.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Date Saved")
}
