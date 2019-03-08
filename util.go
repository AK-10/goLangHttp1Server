package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"time"
	"path/filepath"
	"strings"
)

contentTypeMap := map[string]string{
	"html": "text/html",
	"htm": "text/html",
	"txt": "text/plain",
	"css": "text/css",
	"png": "image/png",
	"jpg": "image/jpeg",
	"jpeg": "image/jpeg",
	"gif": "image/gif",
}

func readLines(bytes []byte) []string {
	strSlice := make([]string, 0)
	start := 0
	for i := 0; bytes[i] != 0 ; i++ {
		if bytes[i] == 10 {
			// println(string(bytes[start:i]))
			strSlice = append(strSlice, string(bytes[start:i]))
			start = i + 1
		}
	}
	return strSlice
}

func getUTCTime() string {
	return time.Now().Format("Mon, 2 Jan 2006 15:04:05 GMT")
}

// return values are body, status, msg, string
func getResponseItem(method string, reqPath string) ([]byte, int, string, string) {
	path := "./views"
	rootPath, _ := filepath.Abs(path)
	var statusCode int
	var message string
	var location string
	// switch {
	// case method == "GET" && reqPath == "/":
	// 	path = path + "/root.html"
	// 	statusCode = 301
	// 	message = "Moved Permanently"
	// 	location = "http://localhost:8080/index"
	// case method == "GET" && reqPath == "/index":
	// 	println("index")
	// 	path = path + "/index.html"
	// 	statusCode = 200
	// 	message = "OK"
	// default:
	// 	// ここで謎のエラーハンドル
	// 	println("default")
	// 	path = rootPath + "/errors/404.html"
	// 	statusCode = 404
	// 	message = "Not Found"
	// }


	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(absPath)

	// ディレクトリトラバーサル検査
	if !strings.HasPrefix(absPath, rootPath) {
		path = rootPath + "/errors/404.html"
		statusCode = 404
		message = "Not Found"
	}

	// ディレクトリの場合indexへredirect
	fileInfo, err := os.Stat(absPath)
	if err != nil {
		log.Fatal(err)
	}
	if fileInfo.IsDir() {
		statusCode = 301
		message = "Moved Permanently"
		location = "http://localhost:8080/index"
	} 

	f, err := os.Open(absPath)
	if err != nil {
		fmt.Println(err)
		statusCode = 500
		message = "Internal Server Error"
		f, _ = os.Open(rootPath + "/views/errors/500.html")
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("error was occured in reading file")
	}

	return buf, statusCode, message, location
}