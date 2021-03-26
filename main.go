package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	programExt := "go"

	if len(os.Args) == 2 && os.Args[1] == "d" {
		programExt = "dart"
		fmt.Println("DART!")
	}

	fmt.Println("input workingDirectory (ex. CF, BOJ ...)")
	var workingDirectory string
	fmt.Scanf("%s", &workingDirectory)
	workingDirectory = strings.ToUpper(workingDirectory)

	var subDirectory string
	fmt.Println("input problemName (ex. 264C, 14264)")
	fmt.Scanf("%s", &subDirectory)

	fmt.Printf("subDirectory = %+v\n", subDirectory)

	dirPath := fmt.Sprintf("./%s/%s", workingDirectory, subDirectory)

	if _, err := os.Stat(dirPath); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(dirPath, 0777)
		}
	}

	problemName := subDirectory
	var functionName string = fmt.Sprintf("%s%s", workingDirectory, problemName)

	// problemNumber, err := strconv.Atoi(problemName)

	// tidy
	Tidy(dirPath)
	// if problemName == "" {
	// 	fmt.Println("THIS IS TIDY!")
	// 	// only tidy!
	// } else {
	// publish
	filePath := fmt.Sprintf("./%s/%s/%s.%s", workingDirectory, subDirectory, problemName, programExt)

	fmt.Printf("filePath = %+v\n", filePath)

	if programExt == "dart" {
		PublishFile(filePath, "./dart_template", functionName)
	} else if programExt == "go" {
		PublishFile(filePath, "./template", functionName)
	}

	testFilePath := fmt.Sprintf("./%s/%s/%s_test.%s", workingDirectory, subDirectory, problemName, programExt)
	fmt.Printf("filePath = %+v\n", testFilePath)

	switch programExt {
	case "go":
		PublishFile(testFilePath, "./template_test", functionName)
		break

	case "dart":
		PublishFile(testFilePath, "./dart_template_test", functionName)
		break
	}

	fmt.Println("command is below!!")
	fmt.Printf("cd ./%+v/%+v\n", workingDirectory, subDirectory)
	fmt.Printf("nv -O %+v %+v\n", filePath, testFilePath)

}

/// tidy main
/// comment redundant main function
func Tidy(dirPath string) error {
	dir, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	fmt.Printf("TIDY! dirPath = %+v\n", dirPath)

	for _, file := range dir {
		fileName := file.Name()

		if strings.Contains(fileName, "_test.go") {
			// test file
			continue
		}

		fmt.Printf("fileName = %+v\n", fileName)

		fullPath := fmt.Sprintf("%s/%s", dirPath, fileName)

		in, err := os.Open(fullPath)
		defer in.Close()

		if err != nil {
			log.Fatalln("can not open the file")
			return err
		}
		content, _ := io.ReadAll(in)
		result_content := strings.ReplaceAll(string(content), `// LETYLETYLETY`, "/*")
		result_content = strings.ReplaceAll(result_content, "// YTELYTELYTEL", "*/")
		// remove os import for convenient
		result_content = strings.ReplaceAll(result_content, `"os"`, "")
		// clear content
		in.Truncate(0)

		err = ioutil.WriteFile(fullPath, []byte(result_content), 0777)
		if err != nil {
			log.Fatalln("can not write the file")
			return err
		}
		fmt.Printf("fullPath = %+v\n /* */", fullPath)
	}
	return nil
}

// YYYY -> file name
// make source file and test file for [problemName]
func PublishFile(dst, src, problemName string) error {
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

	// clear then
	out.Truncate(0)
	// write
	out.WriteString(
		result_str,
	)

	return nil
}
