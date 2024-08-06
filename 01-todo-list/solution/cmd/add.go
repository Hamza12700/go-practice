package cmd

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add todo to the list",
	Run:   addTodos,
}

func addTodos(cmd *cobra.Command, args []string) {
	file, err := os.OpenFile("data/todos.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("failed to read todos:", err)
	}
	contents, err := os.ReadFile("data/todos.csv")
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}
	if !strings.Contains(string(contents), "IsComplete") {
		_, err = file.WriteString("Description,CreatedAt,IsComplete\n")
		if err != nil {
			log.Fatalln("failed to write to file:", err)
		}
	}

	defer file.Close()
	for _, arg := range args {
		writer := csv.NewWriter(file)
		defer writer.Flush()

		data := [][]string{
			{arg, time.Now().Local().Format(time.RFC850), "false"},
		}

		for _, record := range data {
			if err := writer.Write(record); err != nil {
				log.Fatalln("failed to write record", err)
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
