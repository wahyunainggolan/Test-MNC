package main

import (
	"fmt"
	"strings"
)

func convertToFormatRupiah(amount int) string {
	str := fmt.Sprintf("%d", amount)
	var result []string
	count := 0

	for i := len(str) - 1; i >= 0; i-- {
		result = append([]string{string(str[i])}, result...)
		count++
		if count%3 == 0 && i != 0 {
			result = append([]string{"."}, result...)
		}
	}

	return strings.Join(result, "")
}

func calculateChange(totalCost, amountPaid int) interface{} {
	if amountPaid < totalCost {
		return "False, kurang bayar"
	}

	change := amountPaid - totalCost
	roundedChange := (change / 100) * 100

	denominations := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	denominationCount := make([]int, len(denominations))

	remaining := roundedChange
	for i := 0; i < len(denominations); i++ {
		for remaining >= denominations[i] {
			remaining -= denominations[i]
			denominationCount[i]++
		}
	}

	// Display result
	fmt.Printf("Kembalian yang harus diberikan kasir: %s, dibulatkan menjadi %s\n",
		convertToFormatRupiah(change),
		convertToFormatRupiah(roundedChange))
	fmt.Println("\nPecahan uang:")

	for i := 0; i < len(denominations); i++ {
		if denominationCount[i] > 0 {
			unit := "lembar"
			if denominations[i] < 1000 {
				unit = "koin"
			}

			fmt.Printf("%d %s %s\n", denominationCount[i], unit, convertToFormatRupiah(denominations[i]))
		}
	}

	return nil
}

func main() {
	fmt.Println("Test Case 1:")
	calculateChange(700649, 800000)

	fmt.Println("\nTest Case 2:")
	calculateChange(575650, 580000)

	fmt.Println("\nTest Case 3:")
	fmt.Println(calculateChange(657650, 600000)) // Output: False, kurang bayar
}
