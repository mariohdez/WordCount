package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	known_flags := map[string]bool{"-c": true, "-l": true, "-w": true, "-m": true}
	data_to_display := "everything"
	was_from_stdin := false
	file_name := ""

	fp := os.Stdin

	switch num_args := len(os.Args); num_args {
	case 1:
		was_from_stdin = true
		break
	case 2:
		file_name = os.Args[1]
		file, err := os.Open(file_name)

		if err != nil {
			fmt.Fprintf(os.Stderr, "file with name %s could not be opened\n", file_name)
			os.Exit(1)
		}

		fp = file

		break
	case 3:
		if _, ok := known_flags[os.Args[1]]; !ok {
			os.Exit(1)
		}

		data_to_display = os.Args[1]
		file_name = os.Args[2]

		file, err := os.Open(file_name)

		if err != nil {
			fmt.Fprintf(os.Stderr, "File with name %s could not be opened\n", file_name)
			os.Exit(1)
		}

		fp = file
	default:
		os.Exit(1)
	}

	total_bytes, total_lines, total_words := readFile(fp, data_to_display)

	defer fp.Close()

	switch {
	case data_to_display == "everything":
		if was_from_stdin {
			fmt.Printf("%8d%8d%8d\n", total_lines, total_words, total_bytes)
		} else {
			fmt.Printf("%8d%8d%8d %8s\n", total_lines, total_words, total_bytes, file_name)
		}
		break
	case data_to_display == "-c":
		if was_from_stdin {
			fmt.Printf("%8d\n", total_lines)
		} else {
			fmt.Printf("%8d %8s\n", total_bytes, file_name)
		}
		break
	case data_to_display == "-l":
		if was_from_stdin {
			fmt.Printf("%8d %8s\n", total_lines)
		} else {
			fmt.Printf("%8d %8s\n", total_lines, file_name)
		}
		break
	case data_to_display == "-w":
		if was_from_stdin {
			fmt.Printf("%8d %8s\n", total_lines)
		} else {
			fmt.Printf("%8d %8s\n", total_words, file_name)
		}
		break
	}


	os.Exit(0)
}

func readFile(fp *os.File, data_to_display string) (int, int, int) {
	total_bytes := 0
	total_lines := 0
	total_words := 0
	buffer_size := 256
	buffer := make([]byte, buffer_size)
	ended_on_a_char := false
	word_len_so_far := 0

	for {
		bytes_read, err := fp.Read(buffer)

		if err != nil {
			if err == io.EOF {
				break
			}
			os.Exit(1)
		}

		total_bytes += bytes_read

		if bytes_read == 0 {
			break
		}

		for j := bytes_read; j < buffer_size; j++ {
			buffer[j] = 0
		}

		runes := bytes.Runes(buffer)

		N := len(runes)

		for i := 0; i < N; {
			if runes[i] == 0 {
				break
			}

			for i < N && !unicode.IsSpace(runes[i]) {
				i++
				word_len_so_far++
			}

			ended_on_a_char = i == N

			if ended_on_a_char {
				break
			}

			for i < N && unicode.IsSpace(runes[i]) {
				// 10 is the codepoint for the new line character.
				if runes[i] == 10 {
					total_lines += 1
				}

				i++
			}

			if word_len_so_far != 0 {
				total_words += 1
				word_len_so_far = 0
			}
		}
	}

	if word_len_so_far != 0 {
		total_words += 1
	}

	return total_bytes, total_lines, total_words
}

