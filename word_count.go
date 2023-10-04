package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	known_flags := map[string]bool{"-c": true, "-l":true, "-w":true, "-m":true}
	data_to_display := "everything"
	fp := os.Stdin

	switch num_args := len(os.Args); num_args {
	case 1:
		break
	case 2:
		break
	case 3:
		if _, ok := known_flags[os.Args[1]]; !ok {
			os.Exit(1)
		}

		data_to_display = os.Args[1]
		file_name := os.Args[2]

		file, err := os.Open(file_name)

		if err != nil {
			fmt.Fprintf(os.Stderr, "File with name %s could not be opened\n", file_name)
			os.Exit(1)
		}

		fp = file
		break
	default:
		os.Exit(1)
		break
	}

	total_bytes, total_lines, total_words := readFile(fp, data_to_display)


	defer fp.Close()

	fmt.Printf("  %d   %d  %d %s\n", total_lines, total_words, total_bytes, os.Args[2])

	os.Exit(0)
}

func readFile(fp *os.File, data_to_display string) (int, int, int) {
	total_bytes := 0
	total_lines := 0
	total_words := 0
	buffer_size := 64

	buffer := make([]byte, buffer_size)

	for {
		bytes_read, err := fp.Read(buffer)

		if err != nil {
			if (err == io.EOF) {
				break
			}
			os.Exit(1)
		}

		total_bytes += bytes_read

		for i:= 0; i < bytes_read; {
			prev_location := i
			for i < bytes_read && !isWhitespace(buffer[i]) {
				i++
			}


			if prev_location != i {
				total_words += 1
			}

			for i < bytes_read && isWhitespace(buffer[i]) {
				if (buffer[i] == 10) {
					total_lines += 1
				}

				i++
			}
		}
	}

	return total_bytes, total_lines, total_words
}

func isWhitespace(b byte) bool {
	if (b > 8 && b < 14) {
		return true
	}

	if (b == 32) {
		return true
	}

	return false
}

