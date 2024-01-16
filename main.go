package main

import (
	"io"
	"log"
	"os"
)

const (
	obsdianFileName = "obs_note.md"
	logseqFileName  = "logseq_note.md"
	content         = "Это заметка про язык программирования [[Golang|Go]]."
)

func main() {
	// Create note in Obsidian style
	if err := createObsNote(content); err != nil {
		log.Fatal(err)
	}

	// Convert note from Obsidian style and get updated content in Logseq style
	updatedContent, err := convertNote()
	if err != nil {
		log.Fatal(err)
	}

	// Create note in Logseq style
	if err := createLogseqNote(updatedContent); err != nil {
		log.Fatal(err)
	}
}

func createObsNote(content string) error {
	file, err := os.Create(obsdianFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}

func convertNote() ([]byte, error) {
	file, err := os.Open(obsdianFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteContent, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var (
		content            = string(byteContent)
		updatedContent     string
		startLink, endLink string
		isFullLink         bool
	)

	for _, v := range content {
		if startLink == "" && v != '[' && v != '|' && v != ']' {
			updatedContent += string(v)
		}

		if isFullLink && v != ']' {
			continue
		}

		if startLink == "[[" {
			if v == ']' {
				switch endLink {
				case "":
					endLink += string(v)
					updatedContent += string(v)
				case "]":
					endLink += string(v)
					updatedContent += string(v)
				}

				continue
			}

			if v == '|' {
				isFullLink = true
				continue
			}

			updatedContent += string(v)
		}

		if v == '[' {
			switch endLink {
			case "":
				startLink += string(v)
				updatedContent += string(v)
			case "[":
				startLink += string(v)
				updatedContent += string(v)
			}
		}
	}

	return []byte(updatedContent), err
}

func createLogseqNote(content []byte) error {
	file, err := os.Create(logseqFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	return nil
}
