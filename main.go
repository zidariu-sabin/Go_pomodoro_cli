/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/fatih/color"
)

func main() {
	// cmd.Execute()
	// for {
	// 	fmt.Printf("Time Remaining: %.1f", internal.TimeRemaining.Seconds())
	// }
	var workTime int
	var breakTime int
	flag.IntVar(&workTime, "workTime", 5, "duration of the working time")
	flag.IntVar(&breakTime, "breakTime", 5, "duration of the break time")
	flag.Parse()

	// secondsRemaining := workTime * 60
	// minutes := secondsRemaining / 60
	// seconds := secondsRemaining % 60
	var secondsRemaining int
	var minutes int
	var seconds int
	//duration given is the duration after run that the timer will start
	timer := time.NewTimer(time.Second)
	//duration given is the duration between the ticks
	ticker := time.NewTicker(time.Second)

	defer ticker.Stop()
	timerColor := color.New(color.FgCyan).SprintFunc()

	go func() {
		for {
			//select keyword allows for initiating an individual await on each
			//channel operation in the scope
			select {
			case <-timer.C:
				fmt.Println("Work timer for", workTime, " minutes has started!")
				//#
				secondsRemaining = workTime * 60

				minutes = secondsRemaining / 60
				seconds = secondsRemaining % 60
			case <-ticker.C:
				if secondsRemaining > 0 {
					minutes = secondsRemaining / 60
					seconds = secondsRemaining % 60
					fmt.Printf("\rTime Remaining: %s   ", timerColor(fmt.Sprintf("%02d:%02d", minutes, seconds)))
					secondsRemaining--
				}
				if secondsRemaining == 0 {
					fmt.Println("\n Break timer for", breakTime, "minutes has started!")
					//#
					secondsRemaining = breakTime * 60
				}
			}
		}
	}()

	//necessary to read the ticker, time.Sleep also takes time so ticker might only print 4 times
	//#
	time.Sleep(time.Duration(workTime+breakTime) * time.Minute)
	timer.Stop()

}
