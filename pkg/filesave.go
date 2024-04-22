package pkg

import (
	"fmt"
	"os"
)

func FileSave(txt []string) {
	file, err := os.Create("files/result.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	for _, t := range txt {
		file.WriteString(t)
		file.WriteString("\n")
	}
}
