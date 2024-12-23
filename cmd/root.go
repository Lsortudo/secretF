/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"net/http"

	"github.com/Lsortudo/secretF/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// CORS  manual teste
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Adicionando os cabeçalhos necessários para CORS
		c.Header("Access-Control-Allow-Origin", "*")                                    // Permite qualquer origem, ou substitua * por uma origem específica
		c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")                    // Métodos permitidos
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization") // Cabeçalhos permitidos
		c.Header("Access-Control-Allow-Credentials", "true")                            // Permite credenciais (cookies, autenticação, etc.)

		// Se for uma requisição OPTIONS, retorne uma resposta 200
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// Passa para o próximo middleware ou handler
		c.Next()
	}
}

var serverCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		router := gin.Default()
		router.Use(CORSMiddleware())
		// Route for secret santa drawing
		router.POST("/draw", handler.DrawSecretSanta)

		log.Println("Server started on port 8080")
		if err := router.Run(":8080"); err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() error {
	return serverCmd.Execute()
}

/*
// rootCmd represents the base command when called without any subcommands

	var rootCmd = &cobra.Command{
		Use:   "secretF",
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.secretF.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
*/
