/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ReadFileCmd represents the ReadFile command
var ReadFileCmd = &cobra.Command{
	Use:   "ReadFile",
	Short: "Ler um arquivo e sortear os pares",
	Long:  `Essa funcao irá ler um arquivo TXT, sortear os pares do amigo secreto de forma aleatoria e gerar um arquivo TXT com os pares`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ReadFile called")
	},
}

func init() {
	rootCmd.AddCommand(ReadFileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ReadFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ReadFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
