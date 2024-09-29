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

func extractFromFile(f *flo.FileObj, re *regexp.Regexp) []string {
	data := map[string]struct{}{}
	str := f.AsString()
	if re.MatchString(str) {
		data[re.ReplaceAllString(str, "$1")] = struct{}{}
	}
	res := []string{}
	for d := range data {
		res = append(res, d)
	}
	sort.Strings(res)
	return res
}

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage:   %s [extension] [directory] [regex]\n", filepath.Base(os.Args[0]))
		fmt.Printf("Example: %s eml /emails/ '\\bfoo[bar|]\\b'\n", filepath.Base(os.Args[0]))
		return
	}
	matches := map[string]struct{}{}
	ext := "." + os.Args[1]
	re := regexp.MustCompile(`(?ims)` + os.Args[3])
	flo.Dir(os.Args[2]).Each(func(f *flo.FileObj) {
		if !strings.HasSuffix(f.Path(), ext) {
			return
		}
		for _, e := range extractFromFile(f, re) {
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
