package note

import (
	"io"
	"log"
	"os"
)

func Create(filename string, content []byte) error {
	file, err := os.Create(filename)
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

func Open(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteContent, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return byteContent, nil
}
