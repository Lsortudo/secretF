/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// declarar vars pra usar aqui em cima pra ficar + organizado
var filePath string

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

}
