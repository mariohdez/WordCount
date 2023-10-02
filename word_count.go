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
	case 2:
	case 3:
		if _, ok := known_flags[os.Args[1]]; !ok {
			os.Exit(1)
		}

		data_to_display = os.Args[1]
		file_name := os.Args[2]

		file, err := os.Open(file_name)

		if err != nil {
			os.Exit(1)
		}

		fp = file
	default:
		os.Exit(1)
	}

	readFile(fp, data_to_display)

	fmt.Println(data_to_display)

	defer fp.Close()

	os.Exit(0)
}

func readFile(fp *os.File, data_to_display string) {
	total_bytes := 0
	buffer := make([]byte, 64)

	for {
		bytes_read, err := fp.Read(buffer)

		if err != nil {
			if (err == io.EOF) {
				break
			}
			os.Exit(1)
		}

		total_bytes += bytes_read
	}

	fmt.Println(total_bytes)
}

