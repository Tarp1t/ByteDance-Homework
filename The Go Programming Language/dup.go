package main

//查找输入重复行的内容和次数
import (
	"bufio"
	"fmt"
	"os"
)

type LnFile struct {
	count    int
	Filename []string
}

func main() {
	counts := make(map[string]*LnFile)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup:%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%v %d\t%s\n", n.Filename, n.count, line)

		}
	}
}

func countLines(f *os.File, counts map[string]*LnFile) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		key := input.Text()
		_, ok := counts[key]
		if ok {
			counts[key].count++
			counts[key].Filename = append(counts[key].Filename, f.Name())
		} else {
			counts[key] = new(LnFile)
			counts[key].count = 1
			counts[key].Filename = append(counts[key].Filename, f.Name())
		}
	}
}
