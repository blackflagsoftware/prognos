package util

import (
	"bufio"
	"os"
	"strings"
)

func ParseInput() string {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	return s
}
