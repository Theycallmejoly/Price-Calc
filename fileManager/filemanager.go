package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManger struct {
	inputFilePath  string
	outputFilePath string
}

func (fm FileManger) ReadLines() ([]string, error) {
	file, err := os.Open(fm.inputFilePath)
	if err != nil {
		return nil, errors.New("Failed to Read The Lines from the File")
	}
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		file.Close()
		return nil, errors.New("Failed to Open the File")
	}

	file.Close()
	return lines, nil
}

func (fm FileManger) WriteTheResult(data interface{}) error {
	file, err := os.Create(fm.outputFilePath)

	if err != nil {
		return errors.New("Failed to Create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("Failed to Conver Data to JSON.")
	}

	file.Close()
	return nil
}

func New(inputPath string, outputPath string) FileManger {
	return FileManger{
		inputFilePath:  inputPath,
		outputFilePath: outputPath,
	}
}
