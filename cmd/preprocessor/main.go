package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/Fotkurz/rinha-de-backend-2026-go/internal/config"
	"github.com/Fotkurz/rinha-de-backend-2026-go/pkg/utils"
	"github.com/joho/godotenv"
)

type Vector []float32

type Item struct {
	Vector Vector `json:"vector"`
	Label  string `json:"label"`
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("failed to load .env file")
	}
	config.LoadLogger()

	const input = "assets/references.json.gz"
	const output = "assets/references.bin"

	f := utils.ReadEnvOrDefault("DATASET_FILE", input)

	slog.Info("Loading dataset from file", "file", f)
	vectors, err := readVectors(f)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Finished loading vectors", "totalVectors", len(vectors))
}

func readVectors(filename string) ([]Vector, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer gzr.Close()

	dec := json.NewDecoder(gzr)

	// read opening '['
	t, err := dec.Token()
	if err != nil {
		return nil, err
	}
	if delim, ok := t.(json.Delim); !ok || delim != '[' {
		return nil, fmt.Errorf("expected '[', got %v", t)
	}

	var vectors []Vector
	for dec.More() {
		var item Item
		if err := dec.Decode(&item); err != nil {
			return nil, err
		}

		vec := Vector{}
		for _, v := range item.Vector {
			vec = append(vec, v)
		}
		vectors = append(vectors, vec)
	}

	if _, err := dec.Token(); err != nil {
		return nil, err
	}

	return vectors, nil
}
