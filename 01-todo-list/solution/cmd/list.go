package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todos",
	Run:   listTodos,
}

func listTodos(cmd *cobra.Command, args []string) {
	ok, err := cmd.Flags().GetBool("all")
	if err != nil {
		log.Fatalln("failed to parse arg:", err)
	}

	file, err := os.Open("data/todos.csv")
	if os.IsNotExist(err) {
		fmt.Println("The todo list is empty")
		os.Exit(0)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if _, err = reader.Read(); err != nil {
		log.Fatalln("failed to parse csv file:", err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("failed to read csv file")
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Description", "Created At", "Done"})
	for _, record := range records {
		id, desc, createdAt, isComplete := record[0], record[1], record[2], record[3]
		parseBool, err := strconv.ParseBool(isComplete)
		if err != nil {
			log.Fatalln("failed to parse bool:", err)
		}
		if !ok && parseBool {
			continue
		}

		parseTime, err := time.Parse(time.RFC850, createdAt)
		if err != nil {
			log.Fatalln("failed to parse time:", err)
		}
		t.AppendRows([]table.Row{
			{id, desc, timediff.TimeDiff(parseTime), isComplete},
		})
		t.SetStyle(table.StyleLight)
	}
	t.Render()
}

func init() {
	listCmd.Flags().BoolP("all", "a", false, "list all the todos")
	rootCmd.AddCommand(listCmd)
}
