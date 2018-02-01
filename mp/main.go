package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
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

func main() {
	switch os.Args[1] {
	case "durfmt":
		durFmt(os.Args[2])
	default:
		log.Fatalf("command not supported: %s", os.Args[1])
	}
}
