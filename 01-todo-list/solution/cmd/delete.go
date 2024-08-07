package cmd

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a task",
	Run:   deleteTodo,
}

func deleteTodo(cmd *cobra.Command, args []string) {
	file, err := os.Open("data/todos.csv")
	if err != nil {
		log.Fatalln("failed to open file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("failed to readAll csv file:", err)
	}

	defer file.Close()
	var updatedRecords [][]string
	for _, arg := range args {
		taskID, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalln("failed to parse taskID:", err)
		}

		for _, record := range records {
			if record[0] != strconv.FormatInt(int64(taskID), 10) {
				updatedRecords = append(updatedRecords, record)
			}
		}
	}

	writeFile, err := os.Create("data/todos.csv")
	if err != nil {
		log.Fatalln("failed to open file:", err)
	}
	
	writer := csv.NewWriter(writeFile)
	defer writer.Flush()

	for _, record := range updatedRecords {
		if err = writer.Write(record); err != nil {
			log.Fatalln("failed to write updatedRecords:", err)
		}
	}
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
