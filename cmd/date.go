/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "Tells the current date and time.",
	Long:  `Tells the current date and time.`,
	Run: func(cmd *cobra.Command, args []string) {
		currentDate := time.Now().Format("2006-01-02 15:04:05")
		fmt.Println(currentDate)
	},
}

func init() {
	rootCmd.AddCommand(dateCmd)
}
