package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"sync"
)


type Source int
const (
	Syslog Source = iota
	AWS
	WindowsEvents
	Traffic
)

type StellarTA interface {
	Analyzer() []string
}

type Anl struct {
	FileName string
	FL *os.File
	Snr Source
}

// Open A CSV (Remeber to defer anl.Fl.Close() )
func (anl *Anl) OpenCSV() error {
	file,err := os.Open(anl.FileName)
	if err != nil {
		return fmt.Errorf("Error openning CSV: %v",err)
	}
	anl.FL = file
}

func (anl *Anl) Analyzer() []string {
	var tmplts []string
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}
	// Get all column names
	allColumnNames := records[0]
	// Filter out empty columns and get their names
	filteredRecords, removedColumns := removeEmptyColumns(records)
	// Get the column names after removing empty ones
	columnNamesAfter := getColumnsNames(allColumnNames, removedColumns)
	// Print all column names before and after
	fmt.Println("All Column Names:")
	fmt.Println(strings.Join(allColumnNames, ", "))
	fmt.Println("\nColumn Names After Filtering Empty Columns:")
	fmt.Println(strings.Join(columnNamesAfter, ", "))
	// Get the number of columns after removing empty ones
	columnCount := len(filteredRecords[0])
	// Create a WaitGroup to wait for all sorting goroutines to finish
	var wg sync.WaitGroup
	// Define a sorting function for a specific column index
	sortByColumn := func(wg *sync.WaitGroup, colIndex int) {
		defer wg.Done()
		sort.Slice(filteredRecords[1:], func(i, j int) bool {
			return filteredRecords[i+1][colIndex] < filteredRecords[j+1][colIndex]
		})
		// Write the sorted records to a new CSV file
		outFile, err := os.Create(fmt.Sprintf("sorted_column_%d.csv", colIndex))
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer outFile.Close()
		writer := csv.NewWriter(outFile)
		defer writer.Flush()
		// Write the header
		if err := writer.Write(records[0]); err != nil {
			fmt.Println("Error writing header:", err)
			return
		}
		// Write the sorted records
		for _, record := range filteredRecords[1:] {
			if err := writer.Write(record); err != nil {
				fmt.Println("Error writing record:", err)
				return
			}
		}
		fmt.Printf("CSV file sorted by column %d (%s) successfully!\n", colIndex, records[0][colIndex])
	}

	// Iterate through columns and sort concurrently
	wg.Add(columnCount)
	for col := 0; col < columnCount; col++ {
		go sortByColumn(&wg, col)
	}

	// Wait for all sorting goroutines to finish
	wg.Wait()

	// Print names of removed columns
	if len(removedColumns) > 0 {
		fmt.Println("\nColumns Removed due to being empty:")
		fmt.Println(strings.Join(removedColumns, ", "))
	}
}

func removeEmptyColumns(records [][]string) ([][]string, []string) {
	nonEmptyColumns := []int{}
	removedColumns := []string{}
	// Find non-empty columns
	for colIndex := range records[0] {
		notEmpty := false
		for _, row := range records {
			if len(row[colIndex]) > 0 {
				notEmpty = true
				break
			}
		}
		if notEmpty {
			nonEmptyColumns = append(nonEmptyColumns, colIndex)
		} else {
			removedColumns = append(removedColumns, records[0][colIndex])
		}
	}
	// Filter records to include only non-empty columns
	filteredRecords := [][]string{}
	for _, row := range records {
		filteredRow := []string{}
		for _, colIndex := range nonEmptyColumns {
			filteredRow = append(filteredRow, row[colIndex])
		}
		filteredRecords = append(filteredRecords, filteredRow)
	}

	return filteredRecords, removedColumns
}

func getColumnsNames(allColumns []string, removedColumns []string) []string {
	var remainingColumns []string
	for _, col := range allColumns {
		if !contains(removedColumns, col) {
			remainingColumns = append(remainingColumns, col)
		}
	}
	return remainingColumns
}

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
