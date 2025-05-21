package main

import (
	"fmt"
)

// manual function to convert to lower case
func toLower(s string) string {
	var result string

	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			result += string(ch + 32)
		} else {
			result += string(ch)
		}
	}

	return result
}

func findMatchingStrings(n int, inputs []string) interface{} {
	lowerString := make([]string, n)

	for i := 0; i < n; i++ {
		lowerString[i] = toLower(inputs[i])
	}

	type Data struct {
		count   int
		indices []int
	}

	listUniqueString := []string{}
	listStringData := []Data{}

	for i := 0; i < n; i++ {
		found := false

		for j := 0; j < len(listUniqueString); j++ {
			if lowerString[i] == listUniqueString[j] {
				listStringData[j].count++
				listStringData[j].indices = append(listStringData[j].indices, i+1) // save 1-based index
				found = true
				break
			}
		}
		if !found {
			listUniqueString = append(listUniqueString, lowerString[i])
			listStringData = append(listStringData, Data{count: 1, indices: []int{i + 1}})
		}
	}

	firstDuplicateIndex := n + 1
	resultIndices := []int{}

	for i := 0; i < len(listStringData); i++ {
		if listStringData[i].count > 1 {
			if listStringData[i].indices[1] < firstDuplicateIndex {
				firstDuplicateIndex = listStringData[i].indices[1]
				resultIndices = listStringData[i].indices
			}
		}
	}

	if len(resultIndices) > 0 {
		result := ""

		for i := 0; i < len(resultIndices); i++ {
			if i > 0 {
				result += " "
			}
			result += fmt.Sprintf("%d", resultIndices[i])
		}

		return result
	}

	return false
}

func main() {
	// Contoh input 1
	input1 := []string{"abcd", "acbd", "aaab", "acbd"}
	fmt.Println("Test case 1 :", len(input1), input1)
	fmt.Println("Answer :", findMatchingStrings(len(input1), input1), "\n")

	// Contoh input 2
	input2 := []string{
		"Satu", "Sate", "Tujuh", "Tusuk", "Tujuh", "Sate",
		"Bonus", "Tiga", "Puluh", "Tujuh", "Tusuk",
	}
	fmt.Println("Test case 1 :", len(input2), input2)
	fmt.Println("Answer :", findMatchingStrings(len(input2), input2), "\n")

	// Contoh input 3
	input3 := []string{"pisang", "goreng", "enak", "sekali", "rasanya"}
	fmt.Println("Test case 1 :", len(input3), input3)
	fmt.Println("Answer :", findMatchingStrings(len(input3), input3), "\n")
}
