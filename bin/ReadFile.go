// /*
// Copyright © 2024 NAME HERE <EMAIL ADDRESS>
// */
// package cmd

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/spf13/cobra"
// )

// // declarar vars pra usar aqui em cima pra ficar + organizado
// var filePath string

// var ReadFileCmd = &cobra.Command{
// 	Use:   "ReadFile",
// 	Short: "Ler um arquivo e sortear os pares",
// 	Long:  `Essa funcao irá ler um arquivo TXT, sortear os pares do amigo secreto de forma aleatoria e gerar um arquivo TXT com os pares`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		FunctionReadFile()
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(ReadFileCmd)

// 	ReadFileCmd.PersistentFlags().StringVarP(&filePath, "path", "p", "", "path to TXT file")
// 	ReadFileCmd.MarkPersistentFlagRequired("path")

// }

// func FunctionReadFile() {
// 	f, err := os.Open(filePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	// Começaremos a ler o arquivo aqui
// 	scanner := bufio.NewScanner(f)

// 	// Slice para armazenar os nomes
// 	var names []string

// 	for scanner.Scan() {
// 		// Pega a linha atual
// 		line := scanner.Text()
// 		names = append(names, line)
// 	}

// 	// checaar o erro aqui (posteriormente a verificaacao de impar)
// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	// checagem impar e par
// 	if len(names)%2 != 0 {
// 		fmt.Println("A lista contém um número ímpar de nomes. Não dá para separar em pares.")
// 		return
// 	}
// 	fmt.Println("Nomes armazenados:")
// 	for _, name := range names {
// 		fmt.Println(name) // útilizando FMT pra tirar o horario que o print tava botando no console
// 	}
// }
