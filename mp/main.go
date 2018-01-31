package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	if os.Args[1] != "durfmt" {
		log.Fatalf("this tool doesn't support commands other than durfmt, yet")
	}
	dur, err := time.ParseDuration(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	precision := 2
	x := regexp.MustCompile(`\p{N}+\.?\p{N}+`).ReplaceAllStringFunc(dur.String(), func(s string) string {
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
	fmt.Println(x)
}
