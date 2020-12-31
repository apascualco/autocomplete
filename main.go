package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Case insensitive")
	path, err := filepath.Abs("./file.txt")
	if err != nil {
		log.Fatalln("File not found")
	}

	node, err := loadDataFromFile(path)
	if err != nil && err != io.EOF {
		os.Exit(1)
	}
	e := NewSearchEngine(node)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter prefix (empty to exit): ")
		prefix, _ := reader.ReadString('\n')
		if prefix == "\n" {
			break
		}
		f, err := e.FindByPrefix(strings.ToLower(strings.TrimSpace(prefix)))
		if err == nil {
			sort.Sort(StringSorter(f))
			printResults(f)
			fmt.Println("Found ", len(f))
		} else {
			fmt.Println(err)
		}
	}
}

func printResults(found []string) {
	for _, s := range found {
		fmt.Println(s)
	}
}

func loadDataFromFile(path string) (*Node, error) {
	fmt.Println("Loading resources")
	node := NewNode(false)
	file, err := os.Open(path)
	if err != nil {
		return &Node{}, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	loadWords := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("f")
			return &Node{}, err
		}
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			loadWords++
		}
		node.AddWord(line)
		if err != nil {
			break
		}
	}
	fmt.Println("Words load ", loadWords)
	return node, nil
}
