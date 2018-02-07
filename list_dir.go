package main

import (
	"fmt"
	"os"
	"bufio"
	"path/filepath"
	"strconv"
	"strings"
)

func inputs() (string, int64, int64) {

var searchDir string
	const delim = '\n'
	var minsize, maxsize int64
	
	fmt.Println("Enter the directory to be searched: ")
	
	reader := bufio.NewReader(os.Stdin)
	searchDir, err := reader.ReadString(delim)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println("Enter the minimum size to check: ")
	fmt.Scan(&minsize)
	
	fmt.Println("Enter the maximum size to check: ")
	fmt.Scan(&maxsize)
	
	return searchDir,minsize,maxsize
}

func run(searchDir string, minsize int64, maxsize int64) ([]string, error){
	//searchDir = "C:\\Users\\ADITHYA\\Desktop\\Test Folder"
	searchDir = strings.TrimSpace(searchDir)
	count:=0
	fileList := make([]string, 1)
	
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
	//filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
	if !f.IsDir(){
        size := f.Size()
		size = size/1024
		if size >= minsize && size <= maxsize{
		SizeString := strconv.FormatInt(int64(size),10)
		count += 1
		
		fileList = append(fileList,SizeString, path)
		}
	}
	return err
	})	
	
	if e != nil {
		panic(e)
	}
	
	for i, file := range fileList {
		if i > 0 {
			fmt.Println(file)
		}
	}
	
	fmt.Println("Total Files in directory: ", count)
	return fileList, nil
	
}

func write(fileList []string) {

file, err := os.Create("result.txt")
    if err != nil {
        fmt.Println("Cannot create file", err)
    }
    defer file.Close()
	
	for _, line := range fileList {
		fmt.Fprintln(file, line)
	}
	
}

func main() {
	searchDir, minsize, maxsize := inputs()
	fileList, error:= run(searchDir, minsize, maxsize)
	write(fileList)
	fmt.Println(error)
}