package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"goexpert-desafio-stress-tester/tester"
)

var rootCmd = &cobra.Command{
	Use:   "main.go",
	Short: "Uma CLI para teste de stress",
	Long:  `Envia requisições simultâneas a uma URL, configurável com as flags --url, --requests e --concurrency`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")
		tester.Run(url, requests, concurrency)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "URL para o teste")
	rootCmd.Flags().IntP("requests", "r", 1, "Número total de requisições")
	rootCmd.Flags().IntP("concurrency", "c", 1, "Número de chamadas simultâneas")
	rootCmd.MarkFlagRequired("url")
}
