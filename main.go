/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"pomodoro/cmd"
	"pomodoro/internal"
)

func main() {
	cmd.Execute()
	for {
		fmt.Printf("Time Remaining: %.1f", internal.TimeRemaining.Seconds())
	}
}
