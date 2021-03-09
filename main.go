package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	problemName := os.Args[1]
	fmt.Printf("problemName = %+v\n", problemName)

	fmt.Println("input workingDirectory")
	var workingDirectory string
	fmt.Scanf("%s", &workingDirectory)
	workingDirectory = strings.ToUpper(workingDirectory)

	var functionName string = fmt.Sprintf("%s%s", workingDirectory, problemName)

	fmt.Println("input subDirectory")
	var subDirectory string
	fmt.Scanf("%s", &subDirectory)

	fmt.Printf("subDirectory = %+v\n", subDirectory)

	filePath := fmt.Sprintf("./%s/%s/%s.go", workingDirectory, subDirectory, problemName)

	fmt.Printf("filePath = %+v\n", filePath)
	publishFile(filePath, "./template", functionName)

	filePath = fmt.Sprintf("./%s/%s/%s_test.go", workingDirectory, subDirectory, problemName)
	fmt.Printf("filePath = %+v\n", filePath)
	publishFile(filePath, "./template_test", functionName)
}

/// YYYY -> file name
func publishFile(dst, src, problemName string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}

	result, err := io.ReadAll(in)
	result_str := strings.ReplaceAll(string(result), "YYYY", problemName)

	defer in.Close()

	// 사전에 디렉토리를 생성해야함
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	out.WriteString(
		result_str,
	)

	return nil
}

// func main() {
// 	DirectoryName := os.Args[1]

// 	fmt.Printf("firstArg = %+v\n", DirectoryName)

// 	baseOJPath := fmt.Sprintf("boj")
// 	dirPath := fmt.Sprintf("./boj/%s", DirectoryName)
// 	filePath := fmt.Sprintf("./boj/%s/main.go", DirectoryName)

// 	// create base directories
// 	makeDir(baseOJPath)
// 	makeDir(dirPath)

// 	// create file
// 	dest, err := os.Create(filePath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	src, err := os.Open("template")
// 	if err != nil {
// 		panic(err)
// 	}

// 	/// copy from template
// 	io.Copy(dest, src)
// 	if err != nil {
// 		panic(err)
// 	}

// 	/// start repl
// 	var input string
// 	for {
// 		fmt.Scan(&input)
// 		// fmt.Printf("input = %s\n", input)

// 		if input == "exit" || input == "quit" {
// 			fmt.Println("exit")
// 			break
// 		}
// 	}
// }

// func makeDir(pathname string) {
// 	fmt.Printf("[makeDir] pathname = %+v\n", pathname)

// 	if _, err := os.Stat(pathname); os.IsNotExist(err) {
// 		err = os.Mkdir(pathname, 0700)
// 		if err != nil {
// 			panic(err)
// 		}
// 	} else {
// 		fmt.Println("already exist")
// 	}

// 	fmt.Println("====everything goes well====")
// 	fmt.Printf("with %+v\n", pathname)
// }
