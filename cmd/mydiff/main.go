package main

import (
	"fmt"
	"os"

	"github.com/yakkay/go-exercise/usecases/diff"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
	}
}
func main() {
	equal, err := diff.CompareFiles(os.Getenv("PATH1"), os.Getenv("FILE1"), os.Getenv("PATH2"), os.Getenv("FILE2"))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println()
	if equal {
		fmt.Println("The files are equal")
	}
	if !equal {
		fmt.Println("The files are different")
	}
}
