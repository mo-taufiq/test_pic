package main

import (
	"fmt"
	"sort"
	"strings"
)

func isItLessThan2DuplicatesBooksID(sortedBooksArr []string, bookID string) bool {
	count := 0
	for _, a := range sortedBooksArr {
		if a[:2] == bookID[:2] {
			count++
		}
	}
	if count == 2 { // if there 2 duplicates book id
		return false
	}
	return true
}

func sortBooksByHeight(unsortedBooksHeight []string) []string {
	unsortedBooksHeightMap := []map[string]interface{}{}
	sortedHeightArr := []string{}
	for _, a := range unsortedBooksHeight {
		height := a[2:]
		unsortedBooksHeightMap = append(unsortedBooksHeightMap, map[string]interface{}{
			height: a,
		})
		sortedHeightArr = append(sortedHeightArr, height)
	}

	sort.Slice(sortedHeightArr, func(i, j int) bool {
		return sortedHeightArr[i] > sortedHeightArr[j]
	})

	sortedBookHeightArr := []string{}
	for _, a := range sortedHeightArr {
		for _, b := range unsortedBooksHeightMap {
			if b[a] != nil {
				x, _ := b[a].(string)
				if isItLessThan2DuplicatesBooksID(sortedBookHeightArr, x) {
					sortedBookHeightArr = append(sortedBookHeightArr, x)
				}
			}
		}
	}
	return sortedBookHeightArr
}

func sortBooks(str string) string {
	// set sort rules
	orderByCategory := [10]string{"6", "7", "0", "9", "4", "8", "1", "2", "5", "3"}
	// conver string str to array
	arrBooksProvided := strings.Split(str, " ")
	sortedBooksBasedOnRules := []string{}
	for _, a := range orderByCategory {
		unsortedHeightBooks := []string{}
		for _, b := range arrBooksProvided {
			if a == strings.Split(b, "")[0] {
				unsortedHeightBooks = append(unsortedHeightBooks, b)
			}
		}
		sortedBooksBasedOnRules = append(sortedBooksBasedOnRules, sortBooksByHeight(unsortedHeightBooks)...)
	}
	return strings.Join(sortedBooksBasedOnRules, " ")
}

func main() {
	fmt.Println(sortBooks("3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14"))
}
