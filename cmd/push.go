/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		pushCmd := exec.Command("git", "push", "origin", "main")
		stdin, err := pushCmd.StdinPipe()
		if err != nil {
			log.Printf("error while pushing: %v", err)
		}
		go func() {
			defer stdin.Close()
			io.WriteString(stdin, "rishabhraghwendra18")
		}()
		out, err := pushCmd.CombinedOutput()
		// if err := pullCmd.Start(); err != nil {
		// 	log.Printf("error while start: %v", err)
		// }
		if err != nil {
			log.Printf("error while runing: %v", err)
		}

		log.Printf("pull called output: %s", out)
		fmt.Println("done")
		fmt.Println("push called")
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
