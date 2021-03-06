package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
	"strings"
)

type freq struct {
	str   string
	count int
}

func Top10(input string) []string {
	result := []freq{}

	input = strings.ReplaceAll(input, "\n", " ") // Convert NEW_LINE to SPACE
	input = strings.ReplaceAll(input, "\t", " ") // Convert TAB to SPACE

	regexp := regexp.MustCompile(" +")
	sliceToAnalize := regexp.Split(input, -1) // Convert spaces sequence to one SPACE

	sort.Slice(sliceToAnalize, func(i, j int) bool { return sliceToAnalize[i] < sliceToAnalize[j] }) // Sort slice to set correct words sequence

	// Get counts of same words
	wordCount := 1
	for i := 1; i < len(sliceToAnalize); i++ {
		if sliceToAnalize[i-1] == sliceToAnalize[i] {
			wordCount++
		} else {
			result = append(result, freq{sliceToAnalize[i-1], wordCount})
			wordCount = 1
		}
	}
	result = append(result, freq{sliceToAnalize[len(sliceToAnalize)-1], wordCount})

	// Sort slice to get frequently appearing words in the top
	sort.Slice(result, func(i, j int) bool { return result[i].count > result[j].count })

	// Get top 10
	var ret []string
	for i := 0; i < len(result) && i < 10; i++ {
		if result[i].str == "" {
			continue
		}
		ret = append(ret, result[i].str)
	}

	return ret
}
