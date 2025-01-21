package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func modulo4WithZeroToFour(value byte) int {
	mod := (value - '0') % 4
	if mod == 0 {
		mod = 4
	}
	return int(mod)
}

func processFileContent(content string) string {
	if len(content) < 8 {
		return ""
	}

	output := "5"
	for i := 1; i <= 7; i++ {
		output += strconv.Itoa(modulo4WithZeroToFour(content[i]))
	}
	output += strconv.Itoa(modulo4WithZeroToFour(content[7]))
	output += strconv.Itoa(modulo4WithZeroToFour(content[6]))
	output += strconv.Itoa(modulo4WithZeroToFour(content[5]))
	output += strconv.Itoa(modulo4WithZeroToFour(content[4]))

	return output
}

func main() {
	var dirPath string
	fmt.Print("Voer een map in met .akb bestanden: ")
	fmt.Scanln(&dirPath)

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Fout bij het lezen van de map:", err)
		os.Exit(1)
	}

	for _, entry := range entries {
		if filepath.Ext(entry.Name()) == ".akb" {
			filePath := filepath.Join(dirPath, entry.Name())
			//fmt.Printf("Verwerken bestand: %s\n", filePath)

			contentBytes, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Fout bij het lezen van bestand:", entry.Name(), err)
				continue
			}

			// Verwijder BOM indien aanwezig
			if len(contentBytes) >= 3 && contentBytes[0] == 0xEF && contentBytes[1] == 0xBB && contentBytes[2] == 0xBF {
				contentBytes = contentBytes[3:] // Skip BOM
			}

			content := strings.TrimSpace(string(contentBytes))
			//fmt.Printf("Ruwe bytes van bestand (%s): %v\n", entry.Name(), contentBytes)
			//fmt.Printf("Stringweergave van bestand (%s): '%s'\n", entry.Name(), string(contentBytes))

			output := processFileContent(content)
			if output != "" {
				//fmt.Printf("Output voor %s: %s\n", entry.Name(), output)
				fmt.Printf("%s:1:%s\n", strings.TrimSuffix(entry.Name(), ".akb"), output)
			} else {
				//fmt.Printf("Ongeldige invoer in bestand: %s\n", entry.Name())
				fmt.Printf("%s:1:Ongeldige invoer\n", strings.TrimSuffix(entry.Name(), ".akb"))
			}
		}
	}
}
