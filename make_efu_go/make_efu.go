package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
)

const (
	WindowsTick    = 10000000
	SecToUnixEpoch = 11644473600
	AttrDirectory  = 0x10
	AttrNormal     = 0x80
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <local directory> [<regex1> <regex2>]\n", os.Args[0])
		os.Exit(1)
	}

	rootPath := os.Args[1]
	var re *regexp.Regexp
	var repl string

	if len(os.Args) > 3 {
		pattern := os.Args[2]
		repl = os.Args[3]

		if pattern != "" {
			pattern = strings.ReplaceAll(pattern, "/", `\\`)
			var err error
			re, err = regexp.Compile(pattern)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid regex pattern: %v\n", err)
				os.Exit(1)
			}
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fprint(writer, "Filename,Size,Date Modified,Date Created,Attributes\r\n")

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error accessing path %q: %v\n", path, err)
			return nil
		}

		windowsPath := strings.ReplaceAll(path, "/", "\\")

		if re != nil {
			windowsPath = re.ReplaceAllString(windowsPath, repl)
		}

		size := info.Size()
		mtime := toWindowsTick(info.ModTime().Unix())
		
		ctime := mtime
		if stat, ok := info.Sys().(*syscall.Stat_t); ok {
			// FIXED: Explicitly convert int32 to int64 for ARM compatibility
			ctime = toWindowsTick(int64(stat.Ctim.Sec))
		}

		attr := AttrNormal
		if info.IsDir() {
			attr = AttrDirectory
		}

		fmt.Fprintf(writer, "\"%s\",%d,%d,%d,%d\r\n", 
			windowsPath, size, mtime, ctime, attr)

		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error walking the path: %v\n", err)
	}
}

func toWindowsTick(unixSeconds int64) int64 {
	return WindowsTick * (unixSeconds + SecToUnixEpoch)
}
