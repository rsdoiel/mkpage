package mkpage

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Codesnip(in io.Reader, out io.Writer) error {
	var (
		inBlock bool
	)
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "```") {
			if inBlock {
				inBlock = false
			} else {
				inBlock = true
				continue
			}
		}
		if inBlock {
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
