package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/toxyl/flo"
)

func extractFromFile(f *flo.FileObj, re *regexp.Regexp, repl string) []string {
	data := map[string]struct{}{}
	str := f.AsString()
	matches := re.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		if len(match) > 1 {
			data[re.ReplaceAllString(match[0], repl)] = struct{}{}
		}
	}

	res := []string{}
	for d := range data {
		res = append(res, d)
	}
	sort.Strings(res)
	return res
}

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Usage:   %s [extension] [directory] [regex] <replace>\n", filepath.Base(os.Args[0]))
		fmt.Printf("Example: %s eml /emails/ '\\b(foo)([bar|])\\b' '$2+$1'\n", filepath.Base(os.Args[0]))
		return
	}
	matches := map[string]struct{}{}
	ext := "." + os.Args[1]
	repl := "$1"
	if len(os.Args) == 5 {
		repl = os.Args[4]
	}
	re := regexp.MustCompile(`(?ims)` + os.Args[3])
	flo.Dir(os.Args[2]).Each(func(f *flo.FileObj) {
		if strings.TrimSpace(ext) != "." && !strings.HasSuffix(f.Path(), ext) {
			return
		}
		for _, e := range extractFromFile(f, re, repl) {
			matches[e] = struct{}{}
		}
	}, nil)
	res := []string{}
	for m := range matches {
		res = append(res, m)
	}
	sort.Strings(res)
	fmt.Println(strings.Join(res, "\n"))
}
