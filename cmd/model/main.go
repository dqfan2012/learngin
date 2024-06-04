package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/joho/godotenv"
)

const modelTemplate = `package {{ .PackageName }}

import "gorm.io/gorm"

type {{ .ModelName }} struct {
    gorm.Model
    // Add your fields here
}
`

type ModelData struct {
	PackageName string
	ModelName   string
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	modelName := flag.String("name", "", "Name of the model to create (e.g. Blog/Post)")
	flag.Parse()

	if *modelName == "" {
		log.Fatal("Model name is required")
	}

	createModel(*modelName)
}

func createModel(modelName string) {
	parts := strings.Split(modelName, "/")
	if len(parts) < 2 {
		log.Fatal("Model name must be in the format 'PackageName/ModelName'")
	}

	packageName := strings.ToLower(parts[0])
	modelName = parts[1]
	modelFileName := strings.ToLower(modelName) + ".go"

	// Create the directory if it doesn't exist
	dirPath := filepath.Join("internal", "models", packageName)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	// Create the model file
	filePath := filepath.Join(dirPath, modelFileName)
	if _, err := os.Stat(filePath); err == nil {
		log.Fatalf("Model file already exists: %s", filePath)
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create model file: %v", err)
	}
	defer file.Close()

	// Generate the model file content
	data := ModelData{
		PackageName: packageName,
		ModelName:   modelName,
	}
	tmpl, err := template.New("model").Parse(modelTemplate)
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}
	if err := tmpl.Execute(file, data); err != nil {
		log.Fatalf("Failed to write to model file: %v", err)
	}

	fmt.Printf("Created model: %s\n", filePath)
}
