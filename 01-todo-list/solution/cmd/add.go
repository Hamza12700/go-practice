package cmd

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
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

	haveDate := true
	if !strings.Contains(string(contents), "ID") {
		haveDate = false
		_, err = file.WriteString("ID,Description,CreatedAt,IsComplete\n")
		if err != nil {
			log.Fatalln("failed to write to file:", err)
		}
	}
	writer := csv.NewWriter(file)
	defer func() {
		file.Close()
		writer.Flush()
	}()

	id := 1
	if haveDate {
		csvFile, err := os.Open("data/todos.csv")
		if err != nil {
			log.Fatalln("failed to open file for reading:", err)
		}
		defer csvFile.Close()

		fileReader := csv.NewReader(csvFile)
		_, err = fileReader.Read()
		if err != nil {
			log.Fatalln("failed to read csv file:", err)
		}

		records, err := fileReader.ReadAll()
		if err != nil {
			log.Fatalln("failed to readAll csv file:", err)
		}
		for _, record := range records {
			recordID, err := strconv.Atoi(record[0])
			if err != nil {
				log.Fatalln("failed to parse int")
			}
			if recordID >= id {
				id += 1
			}
		}
	}

	for _, arg := range args {
		data := [][]string{
			{strconv.FormatInt(int64(id), 10), arg, time.Now().Local().Format(time.RFC850), "false"},
		}

		for _, record := range data {
			if err := writer.Write(record); err != nil {
				log.Fatalln("failed to write record", err)
			}
		}
		id += 1
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatalln("failed to flush writer", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
