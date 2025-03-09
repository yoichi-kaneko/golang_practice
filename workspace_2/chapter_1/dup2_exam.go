package main

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/slices"
	"os"
	"strings"
)

func main() {
	text_counts := make(map[string]int)
	file_counts := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, text_counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, text_counts)

			f2, err := os.Open(arg)

			countFiles(f2, arg, file_counts)
			f.Close()
			f2.Close()
		}
	}
	
	for line, n := range text_counts {
		if n > 1 {
			if _, ok := file_counts[line]; ok {
				fmt.Printf("%d\t%s\t%s\n", n, line, strings.Join(file_counts[line], ", "))
			} else {
				fmt.Printf("%d\t%s\n", n, line)	
			}
		}
	}
}

func countLines(f *os.File, text_counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
    	text_counts[input.Text()]++
    }
}

func countFiles(f *os.File, arg string, file_counts map[string][]string) {
    input := bufio.NewScanner(f)
	text_slice := []string{}

    for input.Scan() {
    	if _, ok := file_counts[input.Text()]; ok {
		    text_slice = file_counts[input.Text()]
		} else {
			text_slice = []string{}
		}
    	
    	if !slices.Contains(text_slice, arg) {
		    file_counts[input.Text()] = append(text_slice, arg)
		}
    }
}
