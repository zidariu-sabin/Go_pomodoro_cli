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
	var workTime int
	var breakTime int
	var rounds int
	flag.IntVar(&workTime, "w", 5, "duration of the working time")
	flag.IntVar(&breakTime, "b", 5, "duration of the break time")
	flag.IntVar(&rounds, "r", 4, "number of rounds the user wants the timer to run for")
	flag.Parse()

	secondsRemaining := workTime * 60
	minutes := secondsRemaining / 60
	seconds := secondsRemaining % 60
	//duration given is the duration between the ticks
	ticker := time.NewTicker(time.Second)

	isWorkPeriod := true
	workColor := color.New(color.FgCyan).SprintFunc()
	breakColor := color.New(color.FgHiGreen).SprintFunc()
	done := make(chan struct{})
	go func() {
		roundsRemaining := rounds
		for roundsRemaining >= 0 {
			if isWorkPeriod {
				secondsRemaining = workTime * 60
			} else {
				secondsRemaining = breakTime * 60
			}
			for secondsRemaining > 0 {
				minutes = secondsRemaining / 60
				seconds = secondsRemaining % 60
				if isWorkPeriod {
					fmt.Printf("\r Round %v/%v %s timer for %v minutes has started!  Time Remaining: %s   ",
						4-roundsRemaining,
						rounds,
						workColor("Work"),
						workTime,
						workColor(fmt.Sprintf("%02d:%02d", minutes, seconds)),
					)
				} else {
					fmt.Printf("\r Round %v/%v %s timer for %v minutes has started!  Time Remaining: %s   ",
						4-roundsRemaining,
						rounds,
						breakColor("Break"),
						breakTime,
						breakColor(fmt.Sprintf("%02d:%02d", minutes, seconds)),
					)
				}
				<-ticker.C
				secondsRemaining--
			}
			isWorkPeriod = !isWorkPeriod
			//a break period has ended, a round has elapsed
			if isWorkPeriod {
				roundsRemaining--
			}
		}
		done <- struct{}{}
	}()

	<-done
	ticker.Stop()
}
