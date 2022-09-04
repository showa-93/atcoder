package main

import (
	"bufio"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	line := s.Text()

	if line[0] == '1' {
		write("No")
		return
	}

	fallColumns := make([]bool, 7)
	fallColumns[0] = line[6] == '0'
	fallColumns[1] = line[3] == '0'
	fallColumns[2] = line[1] == '0' && line[7] == '0'
	fallColumns[3] = line[0] == '0' && line[4] == '0'
	fallColumns[4] = line[2] == '0' && line[8] == '0'
	fallColumns[5] = line[5] == '0'
	fallColumns[6] = line[9] == '0'

	bef := !fallColumns[0]
	for i := 1; i < 6; i++ {
		if !fallColumns[i] {
			bef = true
			continue
		}

		if bef {
			for j := i + 1; j < 7; j++ {
				if !fallColumns[j] {
					write("Yes")
					return
				}
			}
		}
	}

	write("No")
}

func write(s string) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	w.WriteString(s + "\n")
}
