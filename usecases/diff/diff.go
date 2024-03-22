package diff

import (
	"fmt"

	"github.com/yakkay/go-exercise/adapters/files"
)

func CompareFiles(path1, file1, path2, file2 string) (equal bool, err error) {
	file1Hashes, err := files.ReadFile(path1, file1, 100)
	if err != nil {
		fmt.Println(err.Error())
	}
	file2Hashes, err := files.ReadFile(path2, file2, 100)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println()
	fmt.Println("DELTAS")
	fmt.Println()
	equal = true
	for i := 0; i < len(file1Hashes) || i < len(file2Hashes); i++ {
		var h1, h2 string
		if i < len(file1Hashes) {
			h1 = file1Hashes[i]
		}
		if i < len(file2Hashes) {
			h2 = file2Hashes[i]
		}
		if h1 != h2 {
			fmt.Printf("Chunk number %d is different. ", i)
			if h1 == "" {
				fmt.Printf("Hash %s was added\n", h2)
			}
			if h2 == "" {
				fmt.Printf("Hash %s was removed\n", h2)
			}
			if h1 != "" && h2 != "" {
				fmt.Printf("Hash %s was modified\n", h1)
			}
			equal = false
		}
	}

	return
}
