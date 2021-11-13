package terminal

import (
	"bytes"
	"fmt"
	"strings"
)

func FormatBlockMessage(format string, msg string) string {
	var buf bytes.Buffer

	width, _ := GetSize()

	lines, maxLen := splitsBlockLines(msg, width-4) // 2 spaces on the left, 2 on the right
	fullPadding := strings.Repeat(" ", maxLen+4)

	buf.WriteString(fmt.Sprintf("<%s>", format))
	buf.WriteString(fullPadding)
	buf.WriteString("</>\n")
	for _, line := range lines {
		buf.WriteString(fmt.Sprintf("<%s>  ", format))
		lenLine, _ := Stdout.GetFormatter().Format([]byte(line), &buf)
		if n := maxLen - lenLine; n >= 0 {
			buf.WriteString(strings.Repeat(" ", n))
		}
		buf.WriteString("  </>\n")
	}
	buf.WriteString(fmt.Sprintf("<%s>", format))
	buf.WriteString(fullPadding)
	buf.WriteString("</>\n")

	return buf.String()
}

func splitsBlockLines(msg string, width int) ([]string, int) {
	lines := []string{}
	maxLen := 0
	// this can happen in headless mode, like running tests for example
	if width <= 0 {
		width = 80
	}

	for _, line := range strings.Split(msg, "\n") {
		line = strings.Replace(line, "\t", "        ", -1)
		lastLinePos := 0
		inAnOpeningTag := false
		inAClosingTag := false
		inATagBody := false
		length := 0
		var lastChar rune
		for pos, char := range line {
			if char == '<' && lastChar != '\\' {
				if len(line) > pos+1 && line[pos+1] == '/' {
					inAClosingTag = true
					inATagBody = false
				} else {
					inAnOpeningTag = true
				}
			}

			if !inAClosingTag && !inAnOpeningTag {
				length += 1
			}

			if char == '>' && lastChar != '\\' {
				if inAnOpeningTag {
					inAnOpeningTag = false
					inATagBody = true
				} else {
					inAClosingTag = false
					inATagBody = false
				}
			}

			if length >= width && !inAClosingTag && !inAnOpeningTag && !inATagBody {
				lines = append(lines, line[lastLinePos:pos+1])
				maxLen = width
				lastLinePos = pos + 1
				length = 0
			}

			lastChar = char
		}

		if lastLinePos < len(line) {
			lines = append(lines, line[lastLinePos:])
			if length > maxLen {
				maxLen = length
			}

		}
	}

	return lines, maxLen
}
