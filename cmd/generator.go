package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

type templateData struct {
	Name      string
	TableName string
}

func generatorName(command string) (string, bool) {
	if len(os.Args) < 3 || strings.TrimSpace(os.Args[2]) == "" {
		fmt.Printf("Usage: %s Name\n", command)
		return "", false
	}

	name := strings.TrimSpace(os.Args[2])
	first := []rune(name)
	if len(first) == 0 || !unicode.IsLetter(first[0]) {
		fmt.Println("Name must start with a letter.")
		return "", false
	}
	for _, r := range name {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			fmt.Println("Name may only contain letters, numbers, and underscores.")
			return "", false
		}
	}
	name = strings.ReplaceAll(name, "_", " ")
	words := strings.Fields(name)
	for i, word := range words {
		words[i] = strings.ToUpper(word[:1]) + word[1:]
	}
	return strings.Join(words, ""), true
}

func writeTemplate(templatePath, target string, data templateData) {
	if _, err := os.Stat(target); err == nil {
		fmt.Printf("Skipped (already exists): %s\n", target)
		return
	}

	source, err := os.ReadFile(templatePath)
	if err != nil {
		fmt.Printf("Read template: %v\n", err)
		return
	}

	tmpl, err := template.New(filepath.Base(templatePath)).Parse(string(source))
	if err != nil {
		fmt.Printf("Parse template: %v\n", err)
		return
	}

	var output bytes.Buffer
	if err := tmpl.Execute(&output, data); err != nil {
		fmt.Printf("Render template: %v\n", err)
		return
	}
	if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
		fmt.Printf("Create directory: %v\n", err)
		return
	}
	if err := os.WriteFile(target, output.Bytes(), 0644); err != nil {
		fmt.Printf("Create file: %v\n", err)
		return
	}
	fmt.Printf("Created: %s\n", target)
}

func snakeCase(value string) string {
	var result []rune
	for i, r := range value {
		if unicode.IsUpper(r) && i > 0 {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}

func pluralize(value string) string {
	if strings.HasSuffix(value, "s") {
		return value + "es"
	}
	if strings.HasSuffix(value, "y") {
		return strings.TrimSuffix(value, "y") + "ies"
	}
	return value + "s"
}
