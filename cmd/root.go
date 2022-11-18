/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ansybl",
	Short: "An open source proof-of-stake validator utility package",
	Long: `

     _____                        ___.   .__   
    /  _  \   ____   _________.__.\_ |__ |  |  
   /  /_\  \ /    \ /  ___<   |  | | __ \|  |  
  /    |    \   |  \\___ \ \___  | | \_\ \  |__
  \____|__  /___|  /____  >/ ____| |___  /____/
          \/     \/     \/ \/          \/      

Welcome to Ansybl! This CLI tool currently has one main function.

1) Monitor for Canto validator uptime and alert via PagerDuty on missing blocks

To get started type 'ansybl init'
 `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ansybl-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
