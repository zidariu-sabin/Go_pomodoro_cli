/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"strconv"
	"time"

	"pomodoro/internal"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pomodoro",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a timer",
	Long: ` Start a timer with costumizable options 
	w working duration
	b break duration
	c number f cycles
	example: pomodoro start -w 25 -b 5 -c 1
		 
	`,
	Run: start,
}

func start(cmd *cobra.Command, args []string) {
	workingTime, _ := cmd.Flags().GetInt("workingTime")
	// breakTime, _ := cmd.Flags().GetInt("breakTime")
	// cycles, _ := cmd.Flags().GetInt("cycles")

	internal.NewTimer(time.Duration(workingTime))

	internal.TimeRemaining, _ = time.ParseDuration(strconv.Itoa(workingTime) + "m")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.AddCommand(startCmd)

	startCmd.Flags().IntP("workingTime", "w", 25, "length of the working phase in the pomodoro sequence")
	startCmd.Flags().IntP("breakTime", "b", 5, "length of the break phase in the pomodoro sequence")
	startCmd.Flags().IntP("cycles", "c", 1, "number of pomodoro sequences that are going to follow")

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pomodoro.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
