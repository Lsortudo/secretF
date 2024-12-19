/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// declarar vars pra usar aqui em cima pra ficar + organizado
var filePath string

var ReadFileCmd = &cobra.Command{
	Use:   "ReadFile",
	Short: "Ler um arquivo e sortear os pares",
	Long:  `Essa funcao irá ler um arquivo TXT, sortear os pares do amigo secreto de forma aleatoria e gerar um arquivo TXT com os pares`,
	Run: func(cmd *cobra.Command, args []string) {
		FunctionReadFile()
	},
}

func init() {
	rootCmd.AddCommand(ReadFileCmd)

	ReadFileCmd.PersistentFlags().StringVarP(&filePath, "path", "p", "", "path to TXT file")
	ReadFileCmd.MarkPersistentFlagRequired("path")

}

func FunctionReadFile() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// comecaremos a ler o arquivo aqui
	scanner := bufio.NewScanner(f)

	// slice nomes
	var names []string
	for scanner.Scan() {

		line := scanner.Text()
		names = append(names, line)
	}

	// checaar o erro aqui (posteriormente a verificaacao de impar)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// exibir os nomes para fim de exemplo, se funcionar fazer gerar um TXT de output
	log.Println("Nomes lidos do arquivo:", names)
}
