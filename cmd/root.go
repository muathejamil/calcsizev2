/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"calcsizev2/FileWordCount"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "calcsizev2",
	Short: "Calculate the number of words using go routines",
	Long:  `Calculate the number of words using go routines`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		startTime := time.Now()
		fmt.Printf("start at: %s\n", startTime)
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Errorf("Error in file path flag!")
			os.Exit(1)
		}

		totalWordsInDir := FileWordCount.CountTotalWordInDir(path)
		endTime := time.Now()
		fmt.Printf("finished at: %s\n", endTime)
		fmt.Printf("Total consumed time is: %s in ms \n ", endTime.Sub(startTime).Microseconds())
		log.Infof("Total words in the directory %s: %d", path, totalWordsInDir)

	},
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.calcsizev2.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("path", "p", "C:\\Users\\Lenovo\\GolandProjects\\createfiles\\data", "The path of the directory to generate file on")
	rootCmd.Flags().BoolP("perFile", "f", false, "Reading per file")
	rootCmd.Flags().BoolP("perLine", "l", false, "Reading per line")
}
