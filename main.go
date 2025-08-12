/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// cmd.Execute()
	// for {
	// 	fmt.Printf("Time Remaining: %.1f", internal.TimeRemaining.Seconds())
	// }
	var workTime int
	flag.IntVar(&workTime, "workTime", 5, "duration of the working time")
	flag.Parse()

	secondsRemaining := workTime * 60
	minutes := secondsRemaining / 60
	seconds := secondsRemaining % 60

	//duration given is the duration after run that the timer will start
	timer := time.NewTimer(time.Second)
	//duration given is the duration between the ticks
	ticker := time.NewTicker(time.Second)

	defer ticker.Stop()

	go func() {
		for {
			//select keyword allows for initiating an individual await on each
			//channel operation in the scope
			select {
			case <-timer.C:
				fmt.Println("Timer for", workTime, " minutes has started!")
			case <-ticker.C:
				if secondsRemaining > 0 {
					minutes = secondsRemaining / 60
					seconds = secondsRemaining % 60
					fmt.Printf("Time Remaining: %02d:%02d\n", minutes, seconds)
					secondsRemaining--
				}
				// fmt.Println("Tick at:", t.Local().Minute(), ":", t.Local().Second())
			}
		}
	}()

	//necessary to read the ticker, time.Sleep also takes time so ticker might only print 4 times
	time.Sleep(time.Duration(workTime*60) * time.Second)
	timer.Stop()

}
