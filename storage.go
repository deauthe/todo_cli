package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	fileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{
		fileName: fileName,
	}
}

func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data,"","    ") //will convert the data into json format
	if err != nil {
		return err
	}

	return os.WriteFile(s.fileName,fileData,0644) //will write the json data to the file 

}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.fileName)

	if err !=nil {
		return err
	}
	return json.Unmarshal(fileData,data) //will convert the json fileData and populate it into the data (second argument)
}