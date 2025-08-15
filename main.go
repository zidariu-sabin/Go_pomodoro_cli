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
	flag.IntVar(&workTime, "w", 5, "duration of the working time")
	flag.IntVar(&breakTime, "b", 5, "duration of the break time")
	flag.Parse()

	secondsRemaining := workTime * 60
	minutes := secondsRemaining / 60
	seconds := secondsRemaining % 60
	//duration given is the duration after run that the timer will start
	timer := time.NewTimer(time.Second)
	//duration given is the duration between the ticks
	ticker := time.NewTicker(time.Second)

	defer ticker.Stop()
	isWorkPeriod := true
	workColor := color.New(color.FgCyan).SprintFunc()
	breakColor := color.New(color.FgHiGreen).SprintFunc()

	go func() {
		for range ticker.C {
			if secondsRemaining > 0 {
				minutes = secondsRemaining / 60
				seconds = secondsRemaining % 60
				if isWorkPeriod {
					fmt.Printf("\r %s timer for %v minutes has started!  Time Remaining: %s   ",
						workColor("Work"),
						workTime,
						workColor(fmt.Sprintf("%02d:%02d", minutes, seconds)),
					)
				} else {
					fmt.Printf("\r %s timer for %v minutes has started!  Time Remaining: %s   ",
						breakColor("Break"),
						breakTime,
						breakColor(fmt.Sprintf("%02d:%02d", minutes, seconds)),
					)
				}
				secondsRemaining--
			}
			if secondsRemaining == 0 {
				isWorkPeriod = !isWorkPeriod
				if isWorkPeriod {
					secondsRemaining = workTime * 60
				} else {
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
