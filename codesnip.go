package mkpage

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Codesnip(in io.Reader, out io.Writer, language string) error {
	var (
		inCodeBlock bool
	)
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "```") {
			if inCodeBlock {
				inCodeBlock = false
			} else if strings.HasPrefix(line, "```"+language) {
				inCodeBlock = true
				continue
			}
		}
		if inCodeBlock {
			switch {
			case strings.HasPrefix(line, "    "):
				fmt.Fprintln(out, strings.TrimPrefix(line, "    "))
			case strings.HasPrefix(line, "\t"):
				fmt.Fprintln(out, strings.TrimPrefix(line, "\t"))
			default:
				fmt.Fprintln(out, line)
			}
		}
	}
	return scanner.Err()
}
