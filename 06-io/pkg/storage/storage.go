package storage

import (
	"encoding/json"
	"gosearch_io/pkg/crawler"
	"os"
)

// Save сохраняет данные source в файл name
func Save(source []crawler.Document, name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	byteData, err := json.Marshal(source)
	if err != nil {
		return err
	}
	_, err = f.Write(byteData)
	if err != nil {
		return err
	}
	return nil
}

// Load загружает данные из файла name
func Load(name string) ([]crawler.Document, error){
	byteData, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	docs := make([]crawler.Document, 0)
	err = json.Unmarshal(byteData, &docs)
	if err != nil {
		return nil, err
	}
	return docs, nil
}
