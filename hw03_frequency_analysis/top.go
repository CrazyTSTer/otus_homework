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

	input = strings.Replace(input, "\n", " ", -1) //Convert NEW_LINE to SPACE
	input = strings.Replace(input, "\t", " ", -1) //Convert TAB to SPACE

	regexp := regexp.MustCompile(" +")
	sliceToAnalize := regexp.Split(input, -1) //Convert spaces sequence to one SPACE

	sort.Slice(sliceToAnalize, func(i, j int) bool { return sliceToAnalize[i] < sliceToAnalize[j] })

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
	sort.Slice(result, func(i, j int) bool { return result[i].count > result[j].count })

	var ret []string
	for i := 0; i < len(result) && i < 10; i++ {
		if result[i].str == "" {
			continue
		}
		ret = append(ret, result[i].str)
	}

	return ret
}
