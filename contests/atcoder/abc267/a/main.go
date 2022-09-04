package main

import (
	"bufio"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	dayofweek := s.Text()
	m := map[string]string{
		"Monday":    "5",
		"Tuesday":   "4",
		"Wednesday": "3",
		"Thursday":  "2",
		"Friday":    "1",
	}
	w := bufio.NewWriter(os.Stdout)
	w.WriteString(m[dayofweek] + "\n")
	w.Flush()
}
