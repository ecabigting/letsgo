package main

import (
	"encoding/json"
	"os"
)

type Store[T any] struct {
	FileName string
}

func NewStorage[T any](fileName string) *Store[T] {
	return &Store[T]{FileName: fileName}
}

func (s *Store[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Store[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}
