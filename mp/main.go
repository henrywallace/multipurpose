package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

func durFmt(raw ...string) {
	pat := regexp.MustCompile(`\p{N}+\.?\p{N}+`)
	for _, r := range raw {
		dur, err := time.ParseDuration(r)
		if err != nil {
			log.Fatal(err)
		}
		precision := 2
		pretty := pat.ReplaceAllStringFunc(dur.String(), func(s string) string {
			t := strings.Split(s, ".")
			if len(t) == 1 {
				return s
			}
			left, right := t[0], t[1]
			if len(right) > precision {
				right = right[:precision]
			}
			return left + "." + right
		})
		fmt.Println(pretty)
	}
}

func counter(top int, words ...string) {
	type wordCount struct {
		word  string
		count int
	}
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	var wcs []wordCount
	for word, count := range counts {
		wcs = append(wcs, wordCount{
			word:  word,
			count: count,
		})
	}

	sort.Slice(wcs, func(i, j int) bool {
		return wcs[i].count > wcs[j].count
	})

	for i, wc := range wcs {
		if top > -1 && i >= top {
			break
		}
		fmt.Printf("%-16s  %d\n", wc.word, wc.count)
	}
}

func tokenize(lines ...string) []string {
	var tokens []string
	for _, line := range lines {
		tokens = append(tokens, strings.Fields(line)...)
	}
	return tokens
}

func isPipe() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return true
	}
	return false
}

func scanArgsOrStdin() []string {
	var out []string

	if isPipe() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			out = append(out, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
		if len(os.Args) > 2 {
			for _, word := range os.Args[2:] {
				out = append(out, word)
			}
		}
	}

	return out
}

func main() {
	switch os.Args[1] {
	case "durfmt":
		durFmt(os.Args[2])
	case "counter":
		input := scanArgsOrStdin()
		tokens := tokenize(input...)
		counter(-1, tokens...)
	default:
		log.Fatalf("command not supported: %s", os.Args[1])
	}
}
