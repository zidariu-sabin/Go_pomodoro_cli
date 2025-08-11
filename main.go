/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// cmd.Execute()
	// for {
	// 	fmt.Printf("Time Remaining: %.1f", internal.TimeRemaining.Seconds())
	// }

	// timer := time.NewTimer(5 * time.Second)
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			//select keyword allows for initiating an individual await on each
			//channel operation in the scope
			select {
			// case <- timer.C:
			// 	fmt.Println("Timer has started")
			case t := <-ticker.C:
				fmt.Println("Tick at:", t.Local().Minute(), ":", t.Local().Second())
			}
		}
	}()

	//necessary to read the ticker, time.Sleep also takes time so ticker might only print 4 times
	time.Sleep(5 * time.Second)
	ticker.Stop()

}
