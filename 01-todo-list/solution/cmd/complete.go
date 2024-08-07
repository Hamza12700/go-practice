package cmd

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "make task as complete",
	Run:   completeTodo,
}

func completeTodo(cmd *cobra.Command, args []string) {
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

	for _, record := range records {
		if record[0] == "ID" {
			continue
		}

		recordID, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatalln("failed to parse recordID:", err)
		}

		for _, arg := range args {
			taskID, err := strconv.Atoi(arg)
			if err != nil {
				log.Fatalln("failed to parse id:", err)
			}

			if taskID == recordID {
				record[3] = "true"
			}
		}
	}

	writerFile, err := os.Create("data/todos.csv")
	if err != nil {
		log.Fatalln("failed to read csv file:", err)
	}
	defer writerFile.Close()
	writer := csv.NewWriter(writerFile)

	for _, record := range records {
		if err = writer.Write(record); err != nil {
			log.Fatalln("failed to write data:", err)
		}
	}
	writer.Flush()
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
