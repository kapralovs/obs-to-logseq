package main

import (
	"log"

	"github.com/kapralovs/obs-to-logseq/internal/note"
)

const (
	obsidianFileName = "obs_note.md"
	logseqFileName   = "logseq_note.md"
	content          = "Это заметка про язык программирования [[Golang|Go]]."
)

func main() {
	// Create note in Obsidian style (TODO:move to examples)
	if err := note.Create(obsidianFileName, []byte(content)); err != nil {
		log.Fatal(err)
	}

	// Convert note from Obsidian style and get updated content in Logseq style
	contentBytes, err := note.Open(obsidianFileName)
	if err != nil {
		log.Fatal(err)
	}

	// Convert note from Obsidian style and get updated content in Logseq style
	updatedContent, err := note.ConvertLinks(contentBytes)
	if err != nil {
		log.Fatal(err)
	}

	// Create note in Logseq style
	if err := note.Create(logseqFileName, updatedContent); err != nil {
		log.Fatal(err)
	}
}
