package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

func percent() (err error) {
	files, err := filepath.Glob("*")
	if err != nil {
		return err
	}
	const dirRegex = `^\d{4}$`
	regex := regexp.MustCompile(dirRegex)
	t := time.Now()
	minYear := t
	entries := 0
	for _, file := range files {
		if fi, err := os.Lstat(file); err != nil {
			return err
		} else if fi.IsDir() {
			if regex.MatchString(file) {
				const format = "2006"
				yearTime, err := time.Parse(format, file)
				if err != nil {
					return err
				}
				if yearTime.Year() <= minYear.Year() {
					minYear = yearTime
				}
				const glob = "*.md"
				dirEntries, err := filepath.Glob(file + string(filepath.Separator) + glob)
				if err != nil {
					return err
				}
				entries += len(dirEntries)
			}
		}
	}
	if entries > 0 {
		percent := float64(entries) / math.Ceil(t.Sub(minYear).Hours()/24)
		fmt.Println(percent)
	}
	return
}
