package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	os.Link("old.txt", "new.txt")
	os.Symlink("old.txt", "new_sym.txt")

	fmt.Println(filepath.Match("*.txt", "new_sym.txt"))
}
