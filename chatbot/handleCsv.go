package chatbot

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// ParseCsv reads a CSV file and returns its content as a slice of string slices.
func ParseCsv(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

// FindRowByKeyword searches for rows where the specified column contains the keyword.
func FindRowByKeyword(data [][]string, keyword string, column int) ([]string, bool) {
	keyword = strings.ToLower(keyword)
	for _, row := range data[1:] { // Skip the header row
		if strings.Contains(strings.ToLower(row[column]), strings.ToLower(keyword)) {
			return row, true
		}
	}
	return nil, false
}

// ExtractDataByKeyword extracts data from a row based on a column header keyword.
func ExtractDataByKeyword(row []string, headers []string, keyword string) (string, bool) {
	keyword = strings.ToLower(keyword)
	for i, header := range headers {
		if strings.Contains(strings.ToLower(header), strings.ToLower(keyword)) {
			return row[i], true
		}
	}
	return "", false
}

// ProcessQuery processes user query to fetch relevant data from CSV file.
func ProcessQuery(query string, fileName string) (string, error) {
	query = strings.ToLower(query)
	data, err := ParseCsv(fileName)
	if err != nil {
		return "", err
	}

	if len(data) == 0 {
		return "", fmt.Errorf("CSV file is empty")
	}

	riceKeywords := []string{"Sindano", "Basmat", "BG", "BW", "BG", "BR", "IR", "ITA", "Pishori"}
	riceTypeKeyword := ""
	words := strings.Fields(query)
	for _, value := range riceKeywords {
		for _, word := range words {
			if strings.Contains(strings.ToLower(value), word) {
				riceTypeKeyword = value
			}
		}
	}

	otherKeywords := []string{"Varieties","Attitude","Rainfall","Temperature","Soils","SoilPH","plantingInstructions1","plantingInstructions2","seedRatePerHectare","harvesting","yield","FertilizerNursery","FertilizerApplication Method","Notes"}
	infoKeyword := ""

	for _, value := range otherKeywords {
		for _, word := range words {
			if strings.Contains(strings.ToLower(value), word) {
				infoKeyword = value
			}
		}
	}
	println(infoKeyword)

	headers := data[0]

	if len(words) < 2 {
		return "", fmt.Errorf("query must contain at least two keywords")
	}

	// riceTypeKeyword := words[0] // e.g., "narika 1"
	//infoKeyword := words[1] // e.g., "rain"

	// Find the row matching the rice type
	row, found := FindRowByKeyword(data, riceTypeKeyword, 2) // Assuming rice names are in the first column
	if !found {
		return "", fmt.Errorf("rice type not found")
	}

	// Extract the data based on the second keyword
	result, found := ExtractDataByKeyword(row, headers, infoKeyword)
	if !found {
		return "", fmt.Errorf("info for keyword %s not found", infoKeyword)
	}

	return result, nil
}
