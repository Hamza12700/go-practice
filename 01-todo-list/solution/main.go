package main

import (
	"log"
	"os"
	"todo-cli-app/cmd"
)

func main() {
	_, err := os.Stat("data")
	if os.IsNotExist(err) {
		err = os.Mkdir("data", os.ModePerm)
		if err != nil {
			log.Fatalln("failed to create directory:", err)
		}
	} else if err != nil {
		log.Fatalln("failed to retrive directory info:", err)
	}

	cmd.Execute()
}
