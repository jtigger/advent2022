package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size uint64
}

type Directory struct {
	name    string
	parent  *Directory
	subdirs map[string]*Directory
	files   map[string]*File
}

func (d *Directory) ContentsSize() (total uint64) {
	for _, subdir := range d.subdirs {
		total += subdir.ContentsSize()
	}
	for _, file := range d.files {
		total += file.Size
	}
	return
}

func (d *Directory) Flatten() []*Directory {
	var subdirs []*Directory

	for _, subdir := range d.subdirs {
		subdirs = append(subdirs, subdir)
		subdirs = append(subdirs, subdir.Flatten()...)
	}

	return subdirs
}

func NewDirectory(parent *Directory, name string) *Directory {
	d := Directory{}
	if parent != nil {
		d.parent = parent
	} else {
		d.parent = &d
	}
	d.name = name
	d.subdirs = make(map[string]*Directory)
	d.files = make(map[string]*File)
	return &d
}

type FileSystem struct {
	root *Directory
	cwd  *Directory
}

func NewFileSystem() *FileSystem {
	fs := FileSystem{}
	fs.root = NewDirectory(nil, "/")
	fs.cwd = fs.root
	return &fs
}

type CDCmd struct {
	fs   *FileSystem
	path string
}

func (c CDCmd) Execute() error {
	var dir *Directory
	switch c.path {
	case "/":
		dir = c.fs.root
	case "..":
		// assumes: parent of root dir _is_ root
		dir = c.fs.cwd.parent
	default:
		var found bool
		dir, found = c.fs.cwd.subdirs[c.path]
		if !found {
			return fmt.Errorf("%s: No such directory", c.path)
		}
	}
	c.fs.cwd = dir
	return nil
}

type LSCmd struct {
	fs    *FileSystem
	stdin *bufio.Scanner
}

func (c LSCmd) Execute() error {
	ok := c.stdin.Scan()
	line := c.stdin.Text()
	for ok && line[0] != '$' {
		tokens := strings.Split(line, " ")
		if tokens[0] == "dir" {
			c.fs.cwd.subdirs[tokens[1]] = NewDirectory(c.fs.cwd, tokens[1])
		} else {
			size, err := strconv.ParseUint(tokens[0], 10, 64)
			if err != nil {
				return fmt.Errorf("ls: Expected file size for %s; was %s", tokens[1], tokens[0])
			}
			c.fs.cwd.files[tokens[1]] = &File{tokens[1], size}
		}
		ok = c.stdin.Scan()
		line = c.stdin.Text()
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fs := NewFileSystem()
	ok := scanner.Scan()
	for ok {
		line := scanner.Text()
		if line == "" {
			ok = scanner.Scan()
			continue
		}

		if line[0] != '$' {
			log.Panicf("expected command line input, but was: %s", line)
		}
		tokens := strings.Split(line, " ")
		switch tokens[1] {
		case "cd":
			cmd := CDCmd{fs, tokens[2]}
			err := cmd.Execute()
			if err != nil {
				log.Panicf("cd: %s", err)
			}
			ok = scanner.Scan()
		case "ls":
			var err error
			cmd := LSCmd{fs, scanner}
			err = cmd.Execute()
			if err != nil {
				log.Panicf("ls: %s", err)
			}
		default:
			log.Panicf("unknown command %s", tokens[1])
		}
	}
	dirs := fs.root.Flatten()
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].ContentsSize() < dirs[j].ContentsSize()
	})
	totalSize := uint64(0)
	for _, dir := range dirs {
		if dir.ContentsSize() > 100_000 {
			break
		}
		totalSize += dir.ContentsSize()
	}
	fmt.Printf("sum dirs <= 100,000 = %d\n", totalSize)
}
