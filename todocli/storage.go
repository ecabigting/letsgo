package main

import (
	"encoding/json"
	"os"
)

// Using generics to
// store not just Todos
type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}

func (s *Storage[T]) Save(data T) error {
	// Using MarshalIndent saves the data into json format with 4 spaces
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	// Write the File
	// 0644 - owner can read write to the file, user group members can read, everyone else can read
	return os.WriteFile(s.FileName, fileData, 0644)

}

func (s *Storage[T]) Load(data *T) error {
	// Read the file from the filesystem
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileData, data)
}
