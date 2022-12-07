package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const Threshold = 100000

// This is surely not the simplest approach, but it seemed like it'd be fun to
// build a trie!
type Directory struct {
	Name string
	Parent *Directory
	Children []*Directory
	TotalSize int
}

func (d Directory) FindChild(name string) *Directory {
	for _, child := range d.Children {
		if child.Name == name {
			return child
		}
	}

	return nil
}

func Part1(dir *Directory, parentPath string) int {
	// fmt.Printf("Checking %s%s/ which has a recursive size of %d\n", parentPath, dir.Name, dir.TotalSize)

	matchingSize := 0

	if dir.TotalSize <= Threshold {
		matchingSize += dir.TotalSize
	}

	for _, child := range dir.Children {
		matchingSize += Part1(child, parentPath + dir.Name + "/")
	}

	return matchingSize
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	root := new(Directory)
	current := root

	for scanner.Scan() {
		line := scanner.Text()

		if line == "$ ls" {
			// Nothing to do
		} else if strings.HasPrefix(line, "$") {
			var dir string
			fmt.Sscanf(line, "$ cd %s", &dir)
			if dir == "/" {
				current = root
			} else if dir == ".." {
				current = current.Parent
			} else {
				current = current.FindChild(dir)
			}
		} else if strings.HasPrefix(line, "dir") {
			var dir string
			fmt.Sscanf(line, "dir %s", &dir)
			child := &Directory{Name: dir, Parent: current}
			current.Children = append(current.Children, child)
		} else {
			var name string
			var size int
			fmt.Sscanf(line, "%d %s", &size, &name)

			for d := current; d != nil; d = d.Parent {
				d.TotalSize += size
			}
		}
	}

	fmt.Println(Part1(root, ""))
}
