/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	// "encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Take pull from repo",
	Run: func(cmd *cobra.Command, args []string) {
		pullCmd := exec.Command("git", "remote", "-v")
		stdout, err := pullCmd.CombinedOutput()
		if err != nil {
			log.Printf("error: %v", err)
		}
		// if err := pullCmd.Start(); err != nil {
		// 	log.Printf("error while start: %v", err)
		// }

		log.Printf("pull called output: %s", stdout)
		fmt.Println("done")
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
