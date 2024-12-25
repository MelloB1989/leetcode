package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	generateSpaces := func(n int, builder *strings.Builder) {
		for i := 0; i < n; i++ {
			builder.WriteByte(' ')
		}
	}
	result := []string{}
	pointer := 0
	currentLen := 0
	start := 0

	for pointer < len(words) {
		if currentLen+len(words[pointer])+(pointer-start) > maxWidth {
			num_words := pointer - start
			num_gaps := num_words - 1
			var builder strings.Builder

			if num_gaps == 0 {
				builder.WriteString(words[start])
				generateSpaces(maxWidth-currentLen, &builder)
			} else {
				total_spaces := maxWidth - currentLen
				spaces_per_gap := total_spaces / num_gaps
				extra_spaces := total_spaces % num_gaps

				for i := start; i < pointer; i++ {
					builder.WriteString(words[i])
					if i < pointer-1 {
						spaces := spaces_per_gap
						if extra_spaces > 0 {
							spaces += 1
							extra_spaces -= 1
						}
						generateSpaces(spaces, &builder)
					}
				}
			}

			result = append(result, builder.String())
			start = pointer
			currentLen = 0
		}

		currentLen += len(words[pointer])
		pointer += 1
	}

	var lastLine strings.Builder
	for i := start; i < pointer; i++ {
		lastLine.WriteString(words[i])
		if i < pointer-1 {
			lastLine.WriteByte(' ')
		}
	}
	generateSpaces(maxWidth-lastLine.Len(), &lastLine)
	result = append(result, lastLine.String())

	return result
}

func main() {
	fmt.Println(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
}
